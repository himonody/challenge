import React, { forwardRef, useImperativeHandle, useState } from "react";
import { Modal, Form, Input, InputNumber } from "antd";
import { addSettlementApi, updateSettlementApi, type ChallengeSettlementModel } from "@/api/app/challenge/settlement";
import { message } from "@/hooks/useMessage";

export interface FormModalRef {
	showAddFormModal: () => void;
	showEditFormModal: (id: number, record?: ChallengeSettlementModel) => void;
}

interface Props {
	onSuccess?: () => void;
}

const FormModal = forwardRef<FormModalRef, Props>(({ onSuccess }, ref) => {
	const [open, setOpen] = useState(false);
	const [confirmLoading, setConfirmLoading] = useState(false);
	const [isEdit, setIsEdit] = useState(false);
	const [currentId, setCurrentId] = useState<number>();
	const [form] = Form.useForm<ChallengeSettlementModel>();

	useImperativeHandle(ref, () => ({
		showAddFormModal: () => {
			setIsEdit(false);
			setCurrentId(undefined);
			form.resetFields();
			setOpen(true);
		},
		showEditFormModal: (id: number, record?: ChallengeSettlementModel) => {
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
				await updateSettlementApi(currentId, values);
				message.success("更新成功");
			} else {
				await addSettlementApi(values);
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
			title={isEdit ? "编辑结算" : "新增结算"}
			open={open}
			onOk={handleOk}
			onCancel={() => setOpen(false)}
			confirmLoading={confirmLoading}
			destroyOnClose
		>
			<Form form={form} layout="vertical">
				<Form.Item label="挑战ID" name="challengeId" rules={[{ required: true, message: "请输入挑战ID" }]}>
					<InputNumber min={1} style={{ width: "100%" }} />
				</Form.Item>
				<Form.Item label="用户ID" name="userId" rules={[{ required: true, message: "请输入用户ID" }]}>
					<InputNumber min={1} style={{ width: "100%" }} />
				</Form.Item>
				<Form.Item label="奖励金额" name="reward" rules={[{ required: true, message: "请输入奖励金额" }]}>
					<Input />
				</Form.Item>
			</Form>
		</Modal>
	);
});

FormModal.displayName = "FormModal";

export default FormModal;
