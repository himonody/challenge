import request from "@/utils/request";
import { ReqPage, ResPage } from "@/utils/request/interface";

export interface RiskStrategyCacheModel {
	scene?: string;
	identityType?: string;
	ruleCode?: string;
	windowSeconds?: number;
	threshold?: number;
	action?: string;
	actionValue?: number;
}

export const getRiskStrategyCachePageApi = (params: ReqPage) =>
	request.get<ResPage<RiskStrategyCacheModel>>("/admin-api/v1/risk/strategy/cache", { ...params, pageIndex: params?.current });

export const exportRiskStrategyCacheApi = (query: object) => request.download("/admin-api/v1/risk/strategy/cache/export", query);
