import { ethers } from "hardhat";

// Error type enum
export enum ErrorType {
  VALIDATION_ERROR = "VALIDATION_ERROR",
  DEPLOYMENT_ERROR = "DEPLOYMENT_ERROR",
  CONFIGURATION_ERROR = "CONFIGURATION_ERROR",
  PERMISSION_ERROR = "PERMISSION_ERROR",
}

// Error info interface
export interface ErrorInfo {
  type: ErrorType;
  message: string;
  details?: string;
  suggestion?: string;
}

// Unified error class
export class BridgeError extends Error {
  public readonly type: ErrorType;
  public readonly details?: string;
  public readonly suggestion?: string;

  constructor(errorInfo: ErrorInfo) {
    super(errorInfo.message);
    this.name = "BridgeError";
    this.type = errorInfo.type;
    this.details = errorInfo.details;
    this.suggestion = errorInfo.suggestion;
  }
}

// Validation errors
export class ValidationError extends BridgeError {
  constructor(message: string, details?: string, suggestion?: string) {
    super({
      type: ErrorType.VALIDATION_ERROR,
      message,
      details,
      suggestion,
    });
  }
}

// Deployment errors
export class DeploymentError extends BridgeError {
  constructor(message: string, details?: string, suggestion?: string) {
    super({
      type: ErrorType.DEPLOYMENT_ERROR,
      message,
      details,
      suggestion,
    });
  }
}

// Configuration errors
export class ConfigurationError extends BridgeError {
  constructor(message: string, details?: string, suggestion?: string) {
    super({
      type: ErrorType.CONFIGURATION_ERROR,
      message,
      details,
      suggestion,
    });
  }
}

// Permission errors
export class PermissionError extends BridgeError {
  constructor(message: string, details?: string, suggestion?: string) {
    super({
      type: ErrorType.PERMISSION_ERROR,
      message,
      details,
      suggestion,
    });
  }
}

// Validate address format
export function validateAddress(address: string, name: string): void {
  if (!address || !ethers.isAddress(address)) {
    throw new ValidationError(
      `Invalid ${name} address`,
      `Address: ${address}`,
      `Please provide a valid Ethereum address`
    );
  }
}

// Validate address array
export function validateAddressArray(addresses: string[], name: string): void {
  if (!addresses || addresses.length === 0) {
    throw new ValidationError(
      `Missing ${name} address`,
      `Address count: ${addresses?.length || 0}`,
      `Please provide ${name} address via env var or CLI argument`
    );
  }

  for (let i = 0; i < addresses.length; i++) {
    if (!ethers.isAddress(addresses[i])) {
      throw new ValidationError(
        `Invalid ${name} address`,
        `Index ${i}: ${addresses[i]}`,
        `Please check address #${i + 1} format`
      );
    }
  }
}

// Validate matching array lengths
export function validateArrayLength(
  array1: any[],
  array2: any[],
  name1: string,
  name2: string
): void {
  if (array1.length !== array2.length) {
    throw new ValidationError(
      `${name1} and ${name2} count mismatch`,
      `${name1} count: ${array1.length}, ${name2} count: ${array2.length}`,
      `Please ensure ${name1} and ${name2} have the same length`
    );
  }
}

// Validate numeric parameter
export function validateNumber(
  value: number,
  name: string,
  min?: number,
  max?: number
): void {
  if (isNaN(value) || !isFinite(value)) {
    throw new ValidationError(
      `Invalid ${name} value`,
      `Value: ${value}`,
      `Please provide a valid number`
    );
  }

  if (min !== undefined && value < min) {
    throw new ValidationError(
      `${name} value too small`,
      `Current: ${value}, minimum: ${min}`,
      `Please ensure ${name} is not less than ${min}`
    );
  }

  if (max !== undefined && value > max) {
    throw new ValidationError(
      `${name} value too large`,
      `Current: ${value}, maximum: ${max}`,
      `Please ensure ${name} is not greater than ${max}`
    );
  }
}

// Format error message
export function formatError(error: any): string {
  if (error instanceof BridgeError) {
    let formatted = `[${error.type}] ${error.message}`;

    if (error.details) {
      formatted += `\nDetails: ${error.details}`;
    }

    if (error.suggestion) {
      formatted += `\nSuggestion: ${error.suggestion}`;
    }

    return formatted;
  }

  return error.message || error.toString();
}

// Unified error handling
export function handleError(error: any): never {
  const formattedError = formatError(error);
  console.error("❌ Operation failed:");
  console.error(formattedError);

  if (error instanceof BridgeError) {
    process.exit(1);
  } else {
    console.error("Unknown error type, please check the logs");
    process.exit(1);
  }
}
