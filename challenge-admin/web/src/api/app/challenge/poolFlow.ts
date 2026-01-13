import request from "@/utils/request";
import { ReqPage, ResPage } from "@/utils/request/interface";

export interface ChallengePoolFlowModel {
	id?: number;
	poolId?: number;
	userId?: number;
	amount?: string;
	type?: number;
	createdAt?: string;
}

export const getPoolFlowPageApi = (params: ReqPage) =>
	request.get<ResPage<ChallengePoolFlowModel>>("/admin-api/v1/challenge/pool/flow", {
		...params,
		pageIndex: params?.current
	});

export const exportPoolFlowApi = (query: object) => request.download("/admin-api/v1/challenge/pool/flow/export", query);

export const addPoolFlowApi = (data: object) => request.post<object>("/admin-api/v1/challenge/pool/flow", data);

export const updatePoolFlowApi = (id: number, data: object) =>
	request.put<object>(`/admin-api/v1/challenge/pool/flow/${id}`, data);

export const delPoolFlowApi = (id: number) => request.delete<object>(`/admin-api/v1/challenge/pool/flow/${id}`);
