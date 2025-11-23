import type { ApiResponse } from "@/shared";

export const AlreadyVerifiedCode = "ALREADY_VERIFIED";
export const InvalidCodeCode = "INVALID_VERIFICATION_CODE";
export const VerifiedCode = "USER_VERIFIED";

type VerificationCode = typeof AlreadyVerifiedCode | typeof InvalidCodeCode | typeof VerifiedCode;

export type VerificationResponse = ApiResponse<VerificationCode>;