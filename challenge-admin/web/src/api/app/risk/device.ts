import request from "@/utils/request";
import { ReqPage, ResPage } from "@/utils/request/interface";

export interface RiskDeviceModel {
	id?: number;
	deviceFp?: string;
	userId?: number;
	createdAt?: string;
}

export const getRiskDevicePageApi = (params: ReqPage) =>
	request.get<ResPage<RiskDeviceModel>>("/admin-api/v1/risk/device", { ...params, pageIndex: params?.current });

export const exportRiskDeviceApi = (query: object) => request.download("/admin-api/v1/risk/device/export", query);
