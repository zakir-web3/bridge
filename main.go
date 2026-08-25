package main

import (
	"os"
	"time"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var Version = ""

func main() {
	rootCmd := &cobra.Command{
		Use:           "bridge",
		Version:       Version,
		SilenceErrors: true,
		SilenceUsage:  true,
		PreRunE:       loadConfig,
		RunE:          run,
	}
	if err := rootCmd.Execute(); err != nil {
		rootCmd.PrintErrf("Error: %+v\n", err)
		os.Exit(1)
	}
}

func run(_ *cobra.Command, _ []string) error {
	cfg := new(Config)
	if err := viper.Unmarshal(cfg, viperDecoderOption); err != nil {
		return errors.Wrap(err, "viper unmarshal config")
	}
	cfg.MergeConfig()
	log.Debug().Str("config", cfg.String()).Msg("loaded config")
	if err := cfg.Validate(); err != nil {
		return err
	}
	// init global logger using config
	if err := Init(cfg.LogLevel, cfg.LogFormat); err != nil {
		return err
	}
	return Start(cfg)
}

func loadConfig(cmd *cobra.Command, _ []string) error {
	if err := viper.BindPFlags(cmd.Flags()); err != nil {
		return errors.Wrap(err, "viper bind flags")
	}
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if !errors.As(err, &configFileNotFoundError) {
			return errors.Errorf(
				"viper read config error: %s, file: %s",
				err.Error(),
				viper.ConfigFileUsed(),
			)
		}
	}
	return nil
}

// Init configures the global logger according to provided level and format.
// level: debug|info|warn|error|fatal|panic|trace (case-insensitive)
// format: "json" or "console" (default: console)
func Init(levelStr, format string) error {
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	zerolog.CallerSkipFrameCount = 2

	level, err := zerolog.ParseLevel(levelStr)
	if err != nil {
		return errors.Wrapf(err, "invalid log level: %s", levelStr)
	}
	zerolog.SetGlobalLevel(level)

	// Configure writer by format
	switch format {
	case "json":
		// JSON output to stdout
		// Set time format for JSON
		zerolog.TimeFieldFormat = time.RFC3339
		log.Logger = zerolog.New(os.Stdout).With().Timestamp().Caller().Logger()
	default:
		// Human-friendly console writer
		w := zerolog.NewConsoleWriter(func(w *zerolog.ConsoleWriter) {
			w.Out = os.Stdout
			w.TimeFormat = time.RFC3339
		})
		log.Logger = zerolog.New(w).With().Timestamp().Caller().Logger()
	}
	return nil
}
