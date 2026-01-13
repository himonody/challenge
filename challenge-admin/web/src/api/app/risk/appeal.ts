import request from "@/utils/request";
import { ReqPage, ResPage } from "@/utils/request/interface";

export interface RiskAppealModel {
	id?: number;
	userId?: number;
	status?: number;
	deviceFp?: string;
	riskReason?: string;
	appealType?: number;
	appealReason?: string;
	appealEvidence?: string;
	actionResult?: number;
	reviewerId?: number;
	reviewRemark?: string;
	createdAt?: string;
	reviewedAt?: string;
}

export const getRiskAppealPageApi = (params: ReqPage) =>
	request.get<ResPage<RiskAppealModel>>("/admin-api/v1/risk/appeal", { ...params, pageIndex: params?.current });

export const exportRiskAppealApi = (query: object) => request.download("/admin-api/v1/risk/appeal/export", query);

export const reviewRiskAppealApi = (id: number, data: object) => request.put<object>(`/admin-api/v1/risk/appeal/${id}/review`, data);
