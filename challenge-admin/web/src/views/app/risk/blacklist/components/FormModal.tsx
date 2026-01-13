import React, { forwardRef, useImperativeHandle, useState } from "react";
import { Modal, Form, Input, InputNumber, Select } from "antd";
import { addRiskBlacklistApi, updateRiskBlacklistApi, type RiskBlacklistModel } from "@/api/app/risk/blacklist";
import { message } from "@/hooks/useMessage";

export interface FormModalRef {
	showAddFormModal: () => void;
	showEditFormModal: (id: number, record?: RiskBlacklistModel) => void;
}

interface Props {
	onSuccess?: () => void;
}

const typeOptions = [
	{ label: "IP", value: "ip" },
	{ label: "设备", value: "device" },
	{ label: "国家", value: "country" },
	{ label: "手机", value: "mobile" },
	{ label: "邮箱", value: "email" }
];

const statusOptions = [
	{ label: "生效", value: "1" },
	{ label: "失效", value: "2" }
];

const FormModal = forwardRef<FormModalRef, Props>(({ onSuccess }, ref) => {
	const [open, setOpen] = useState(false);
	const [confirmLoading, setConfirmLoading] = useState(false);
	const [isEdit, setIsEdit] = useState(false);
	const [currentId, setCurrentId] = useState<number>();
	const [form] = Form.useForm<RiskBlacklistModel>();

	useImperativeHandle(ref, () => ({
		showAddFormModal: () => {
			setIsEdit(false);
			setCurrentId(undefined);
			form.resetFields();
			setOpen(true);
		},
		showEditFormModal: (id: number, record?: RiskBlacklistModel) => {
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
			if (isEdit && currentId !== undefined) {
				await updateRiskBlacklistApi(currentId, values);
				message.success("更新成功");
			} else {
				await addRiskBlacklistApi(values);
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
			title={isEdit ? "编辑黑名单" : "新增黑名单"}
			open={open}
			onOk={handleOk}
			onCancel={() => setOpen(false)}
			confirmLoading={confirmLoading}
			destroyOnClose
			width={520}
		>
			<Form form={form} layout="vertical">
				<Form.Item label="类型" name="type" rules={[{ required: true, message: "请选择类型" }]}>
					<Select options={typeOptions} />
				</Form.Item>
				<Form.Item label="命中值" name="value" rules={[{ required: true, message: "请输入命中值" }]}>
					<Input />
				</Form.Item>
				<Form.Item label="风险等级" name="riskLevel" rules={[{ required: true, message: "请输入风险等级" }]}>
					<InputNumber min={0} style={{ width: "100%" }} />
				</Form.Item>
				<Form.Item label="封禁原因" name="reason">
					<Input.TextArea rows={3} />
				</Form.Item>
				<Form.Item label="状态" name="status" rules={[{ required: true, message: "请选择状态" }]}>
					<Select options={statusOptions} />
				</Form.Item>
			</Form>
		</Modal>
	);
});

FormModal.displayName = "FormModal";

export default FormModal;
