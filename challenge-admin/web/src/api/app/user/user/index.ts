import request from "@/utils/request";
import { ReqPage, ResPage } from "@/utils/request/interface";
import { UserLevelModel } from "../user-level";

export interface UserModel {
	id?: number;
	levelId?: number;
	userName?: string;
	nickName?: string;
	trueName?: string;
	money?: string;
	freezeMoney?: string;
	email?: string;
	mobileTitle?: string;
	mobile?: string;
	avatar?: string;
	payPwd?: string;
	payStatus?: string;
	pwd?: string;
	refCode?: string;
	registerAt?: string;
	registerIp?: string;
	lastLoginAt?: string;
	lastLoginIp?: string;
	parentId?: number;
	parentIds?: string;
	treeSort?: string;
	treeSorts?: string;
	treeLeaf?: string;
	treeLevel?: number;
	status?: string;
	remark?: string;
	createBy?: number;
	updateBy?: number;
	createdAt?: Date;
	updatedAt?: Date;
	userLevel?: UserLevelModel;
}

export const getUserPageApi = (params: ReqPage) => {
	return request.get<ResPage<UserModel>>(`/admin-api/v1/app/user/user`, { ...params, pageIndex: params?.current });
};

export const getUserApi = (id: number) => {
	return request.get<UserModel>(`/admin-api/v1/app/user/user/` + id);
};

export const addUserApi = (data: object) => {
	return request.post<object>(`/admin-api/v1/app/user/user`, data);
};

export const updateUserApi = (id: number, data: object) => {
	return request.put<object>("/admin-api/v1/app/user/user/" + id, data);
};

export const delUserApi = (params: number[]) => {
	return request.delete<object>(`/admin-api/v1/app/user/user`, { ids: params });
};

export const exportUserApi = (query: object) => {
	return request.download(`/admin-api/v1/app/user/user/export`, query);
};

// 人工充值
export const rechargeUserApi = (id: number, data: { amount: string | number }) => {
	return request.post(`/admin-api/v1/app/user/user/recharge/${id}`, data);
};

// 人工扣款
export const deductUserApi = (id: number, data: { amount: string | number }) => {
	return request.post(`/admin-api/v1/app/user/user/deduct/${id}`, data);
};

// 重置登录密码
export const resetUserPasswordApi = (id: number) => {
	return request.post(`/admin-api/v1/app/user/user/reset/password/${id}`);
};

// 重置支付密码
export const resetUserPayPasswordApi = (id: number) => {
	return request.post(`/admin-api/v1/app/user/user/reset/pay-password/${id}`);
};

// 支付状态变更
export const updateUserPayStatusApi = (id: number, data: { pay_status: string }) => {
	return request.post(`/admin-api/v1/app/user/user/pay-status/${id}`, data);
};
