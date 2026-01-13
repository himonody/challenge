import request from "@/utils/request";
import { ReqPage, ResPage } from "@/utils/request/interface";

export interface ChallengeCheckinModel {
	id?: number;
	challengeId?: number;
	userId?: number;
	checkinDate?: string;
	checkinTime?: string;
	moodCode?: number;
	moodText?: string;
	contentType?: number;
	status?: number;
	createdAt?: string;
}

export const getCheckinPageApi = (params: ReqPage) =>
	request.get<ResPage<ChallengeCheckinModel>>("/admin-api/v1/challenge/checkin", {
		...params,
		pageIndex: params?.current
	});

export const exportCheckinApi = (query: object) => request.download("/admin-api/v1/challenge/checkin/export", query);

export const addCheckinApi = (data: object) => request.post<object>("/admin-api/v1/challenge/checkin", data);

export const updateCheckinApi = (id: number, data: object) => request.put<object>(`/admin-api/v1/challenge/checkin/${id}`, data);

export const delCheckinApi = (id: number) => request.delete<object>(`/admin-api/v1/challenge/checkin/${id}`);
