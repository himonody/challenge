import request from "@/utils/request";
import { ReqPage, ResPage } from "@/utils/request/interface";

export interface ChallengeCheckinImageModel {
	id?: number;
	checkinId?: number;
	userId?: number;
	imageUrl?: string;
	imageHash?: string;
	sortNo?: number;
	status?: number;
	createdAt?: string;
}

export const getCheckinImagePageApi = (params: ReqPage) =>
	request.get<ResPage<ChallengeCheckinImageModel>>("/admin-api/v1/challenge/checkin/image", {
		...params,
		pageIndex: params?.current
	});

export const exportCheckinImageApi = (query: object) => request.download("/admin-api/v1/challenge/checkin/image/export", query);

export const addCheckinImageApi = (data: object) => request.post<object>("/admin-api/v1/challenge/checkin/image", data);

export const updateCheckinImageApi = (id: number, data: object) =>
	request.put<object>(`/admin-api/v1/challenge/checkin/image/${id}`, data);

export const delCheckinImageApi = (id: number) => request.delete<object>(`/admin-api/v1/challenge/checkin/image/${id}`);
