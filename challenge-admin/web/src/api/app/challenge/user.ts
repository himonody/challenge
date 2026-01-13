import request from "@/utils/request";
import { ReqPage, ResPage } from "@/utils/request/interface";

export interface ChallengeUserModel {
  id?: number;
  userId?: number;
  configId?: number;
  poolId?: number;
  challengeAmount?: string;
  startDate?: number;
  endDate?: number;
  status?: number;
  failReason?: number;
  createdAt?: string;
}

export const getChallengeUserPageApi = (params: ReqPage) =>
  request.get<ResPage<ChallengeUserModel>>("/admin-api/v1/challenge/user", {
    ...params,
    pageIndex: params?.current,
  });

export const exportChallengeUserApi = (query: object) =>
  request.download("/admin-api/v1/challenge/user/export", query);

export const addChallengeUserApi = (data: object) =>
  request.post<object>("/admin-api/v1/challenge/user", data);

export const updateChallengeUserApi = (id: number, data: object) =>
  request.put<object>(`/admin-api/v1/challenge/user/${id}`, data);

export const delChallengeUserApi = (id: number) =>
  request.delete<object>(`/admin-api/v1/challenge/user/${id}`);
