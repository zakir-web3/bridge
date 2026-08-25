import { ethers } from "hardhat";

// 错误类型枚举
export enum ErrorType {
  VALIDATION_ERROR = "VALIDATION_ERROR",
  DEPLOYMENT_ERROR = "DEPLOYMENT_ERROR",
  CONFIGURATION_ERROR = "CONFIGURATION_ERROR",
  PERMISSION_ERROR = "PERMISSION_ERROR",
}

// 错误信息接口
export interface ErrorInfo {
  type: ErrorType;
  message: string;
  details?: string;
  suggestion?: string;
}

// 统一错误类
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

// 验证相关错误
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

// 部署相关错误
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

// 配置相关错误
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

// 权限相关错误
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

// 验证地址格式
export function validateAddress(address: string, name: string): void {
  if (!address || !ethers.isAddress(address)) {
    throw new ValidationError(
      `无效的${name}地址`,
      `地址: ${address}`,
      `请提供有效的以太坊地址格式`
    );
  }
}

// 验证地址数组
export function validateAddressArray(addresses: string[], name: string): void {
  if (!addresses || addresses.length === 0) {
    throw new ValidationError(
      `缺少${name}地址`,
      `地址数量: ${addresses?.length || 0}`,
      `请通过环境变量或命令行参数提供${name}地址`
    );
  }

  for (let i = 0; i < addresses.length; i++) {
    if (!ethers.isAddress(addresses[i])) {
      throw new ValidationError(
        `无效的${name}地址`,
        `索引 ${i}: ${addresses[i]}`,
        `请检查第${i + 1}个地址格式`
      );
    }
  }
}

// 验证数组长度匹配
export function validateArrayLength(
  array1: any[],
  array2: any[],
  name1: string,
  name2: string
): void {
  if (array1.length !== array2.length) {
    throw new ValidationError(
      `${name1}和${name2}数量不匹配`,
      `${name1}数量: ${array1.length}, ${name2}数量: ${array2.length}`,
      `请确保${name1}和${name2}的数量一致`
    );
  }
}

// 验证数值参数
export function validateNumber(
  value: number,
  name: string,
  min?: number,
  max?: number
): void {
  if (isNaN(value) || !isFinite(value)) {
    throw new ValidationError(
      `无效的${name}值`,
      `值: ${value}`,
      `请提供有效的数值`
    );
  }

  if (min !== undefined && value < min) {
    throw new ValidationError(
      `${name}值过小`,
      `当前值: ${value}, 最小值: ${min}`,
      `请确保${name}不小于${min}`
    );
  }

  if (max !== undefined && value > max) {
    throw new ValidationError(
      `${name}值过大`,
      `当前值: ${value}, 最大值: ${max}`,
      `请确保${name}不大于${max}`
    );
  }
}

// 格式化错误信息
export function formatError(error: any): string {
  if (error instanceof BridgeError) {
    let formatted = `[${error.type}] ${error.message}`;

    if (error.details) {
      formatted += `\n详细信息: ${error.details}`;
    }

    if (error.suggestion) {
      formatted += `\n建议: ${error.suggestion}`;
    }

    return formatted;
  }

  return error.message || error.toString();
}

// 统一错误处理
export function handleError(error: any): never {
  const formattedError = formatError(error);
  console.error("❌ 操作失败:");
  console.error(formattedError);

  if (error instanceof BridgeError) {
    process.exit(1);
  } else {
    console.error("未知错误类型，请检查日志");
    process.exit(1);
  }
}
