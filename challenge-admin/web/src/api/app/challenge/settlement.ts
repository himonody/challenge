import request from "@/utils/request";
import { ReqPage, ResPage } from "@/utils/request/interface";

export interface ChallengeSettlementModel {
	id?: number;
	challengeId?: number;
	userId?: number;
	reward?: string;
	createdAt?: string;
}

export const getSettlementPageApi = (params: ReqPage) =>
	request.get<ResPage<ChallengeSettlementModel>>("/admin-api/v1/challenge/settlement", {
		...params,
		pageIndex: params?.current
	});

export const exportSettlementApi = (query: object) => request.download("/admin-api/v1/challenge/settlement/export", query);

export const addSettlementApi = (data: object) => request.post<object>("/admin-api/v1/challenge/settlement", data);

export const updateSettlementApi = (id: number, data: object) =>
	request.put<object>(`/admin-api/v1/challenge/settlement/${id}`, data);

export const delSettlementApi = (id: number) => request.delete<object>(`/admin-api/v1/challenge/settlement/${id}`);
