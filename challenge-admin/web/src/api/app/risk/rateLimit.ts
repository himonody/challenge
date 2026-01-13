import request from "@/utils/request";
import { ReqPage, ResPage } from "@/utils/request/interface";

export interface RiskRateLimitModel {
	id?: number;
	scene?: string;
	identityType?: string;
	identityValue?: string;
	count?: number;
	windowStart?: string;
	windowEnd?: string;
	blocked?: string;
	createdAt?: string;
}

export const getRiskRateLimitPageApi = (params: ReqPage) =>
	request.get<ResPage<RiskRateLimitModel>>("/admin-api/v1/risk/rate-limit", { ...params, pageIndex: params?.current });

export const exportRiskRateLimitApi = (query: object) => request.download("/admin-api/v1/risk/rate-limit/export", query);
