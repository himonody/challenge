import request from "@/utils/request";
import { ReqPage, ResPage } from "@/utils/request/interface";

export interface RiskUserModel {
	userId?: number;
	riskLevel?: number;
	riskScore?: number;
	reason?: string;
	updatedAt?: string;
}

export const getRiskUserPageApi = (params: ReqPage) =>
	request.get<ResPage<RiskUserModel>>("/admin-api/v1/risk/user", { ...params, pageIndex: params?.current });

export const exportRiskUserApi = (query: object) => request.download("/admin-api/v1/risk/user/export", query);
