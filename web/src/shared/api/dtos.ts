export interface ApiResponse<CodeType = string> {
  message?: string;
  error?: string;
  code?: CodeType;
}