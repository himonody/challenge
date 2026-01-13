import request from "@/utils/request";
import { ReqPage, ResPage } from "@/utils/request/interface";

export interface ChallengeCheckinVideoAdModel {
	id?: number;
	checkinId?: number;
	userId?: number;
	adPlatform?: string;
	adUnitId?: string;
	adOrderNo?: string;
	videoDuration?: number;
	watchDuration?: number;
	rewardAmount?: string;
	verifyStatus?: number;
	createdAt?: string;
	verifiedAt?: string;
}

export const getCheckinVideoAdPageApi = (params: ReqPage) =>
	request.get<ResPage<ChallengeCheckinVideoAdModel>>("/admin-api/v1/challenge/checkin/video_ad", {
		...params,
		pageIndex: params?.current
	});

export const exportCheckinVideoAdApi = (query: object) =>
	request.download("/admin-api/v1/challenge/checkin/video_ad/export", query);

export const addCheckinVideoAdApi = (data: object) => request.post<object>("/admin-api/v1/challenge/checkin/video_ad", data);

export const updateCheckinVideoAdApi = (id: number, data: object) =>
	request.put<object>(`/admin-api/v1/challenge/checkin/video_ad/${id}`, data);

export const delCheckinVideoAdApi = (id: number) => request.delete<object>(`/admin-api/v1/challenge/checkin/video_ad/${id}`);
