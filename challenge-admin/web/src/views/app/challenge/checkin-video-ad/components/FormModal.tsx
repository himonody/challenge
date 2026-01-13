import React, { forwardRef, useImperativeHandle, useState } from "react";
import { Modal, Form, Input, InputNumber, Select } from "antd";
import type { ChallengeCheckinVideoAdModel } from "@/api/app/challenge/checkinVideoAd";
import { addCheckinVideoAdApi, updateCheckinVideoAdApi } from "@/api/app/challenge/checkinVideoAd";
import { message } from "@/hooks/useMessage";

export interface FormModalRef {
	showAddFormModal: () => void;
	showEditFormModal: (id: number, record?: ChallengeCheckinVideoAdModel) => void;
}

interface Props {
	onSuccess?: () => void;
}

const verifyOptions = [
	{ label: "待校验", value: 0 },
	{ label: "通过", value: 1 },
	{ label: "拒绝", value: 2 }
];

// eslint-disable-next-line react/display-name
const FormModal = forwardRef<FormModalRef, Props>(({ onSuccess }, ref) => {
	const [open, setOpen] = useState(false);
	const [confirmLoading, setConfirmLoading] = useState(false);
	const [isEdit, setIsEdit] = useState(false);
	const [currentId, setCurrentId] = useState<number>();
	const [form] = Form.useForm<ChallengeCheckinVideoAdModel>();

	useImperativeHandle(ref, () => ({
		showAddFormModal: () => {
			setIsEdit(false);
			setCurrentId(undefined);
			form.resetFields();
			setOpen(true);
		},
		showEditFormModal: (id: number, record?: ChallengeCheckinVideoAdModel) => {
			setIsEdit(true);
			setCurrentId(id);
			if (record) form.setFieldsValue(record);
			setOpen(true);
		}
	}));

	const handleOk = async () => {
		try {
			const values = await form.validateFields();
			setConfirmLoading(true);
			if (isEdit && currentId) {
				await updateCheckinVideoAdApi(currentId, values);
				message.success("更新成功");
			} else {
				await addCheckinVideoAdApi(values);
				message.success("新增成功");
			}
			setOpen(false);
			onSuccess?.();
		} catch (e) {
			// ignore
		} finally {
			setConfirmLoading(false);
		}
	};

	return (
		<Modal
			title={isEdit ? "编辑打卡视频广告" : "新增打卡视频广告"}
			open={open}
			onOk={handleOk}
			onCancel={() => setOpen(false)}
			confirmLoading={confirmLoading}
			destroyOnClose
		>
			<Form form={form} layout="vertical">
				<Form.Item label="打卡ID" name="checkinId" rules={[{ required: true, message: "请输入打卡ID" }]}>
					<InputNumber style={{ width: "100%" }} min={1} />
				</Form.Item>
				<Form.Item label="用户ID" name="userId" rules={[{ required: true, message: "请输入用户ID" }]}>
					<InputNumber style={{ width: "100%" }} min={1} />
				</Form.Item>
				<Form.Item label="广告平台" name="adPlatform" rules={[{ required: true, message: "请输入广告平台" }]}>
					<Input />
				</Form.Item>
				<Form.Item label="广告位ID" name="adUnitId">
					<Input />
				</Form.Item>
				<Form.Item label="广告订单号" name="adOrderNo">
					<Input />
				</Form.Item>
				<Form.Item label="视频时长(秒)" name="videoDuration">
					<InputNumber style={{ width: "100%" }} min={0} />
				</Form.Item>
				<Form.Item label="观看时长(秒)" name="watchDuration">
					<InputNumber style={{ width: "100%" }} min={0} />
				</Form.Item>
				<Form.Item label="收益金额" name="rewardAmount">
					<Input />
				</Form.Item>
				<Form.Item label="校验状态" name="verifyStatus" rules={[{ required: true, message: "请选择校验状态" }]}>
					<Select options={verifyOptions} />
				</Form.Item>
			</Form>
		</Modal>
	);
});

export default FormModal;
