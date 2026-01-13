import request from "@/utils/request";
import { ReqPage, ResPage } from "@/utils/request/interface";

export interface RiskEventModel {
	id?: number;
	userId?: number;
	eventType?: number;
	detail?: string;
	score?: number;
	createdAt?: string;
}

export const getRiskEventPageApi = (params: ReqPage) =>
	request.get<ResPage<RiskEventModel>>("/admin-api/v1/risk/event", { ...params, pageIndex: params?.current });

export const exportRiskEventApi = (query: object) => request.download("/admin-api/v1/risk/event/export", query);
