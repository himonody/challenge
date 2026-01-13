import React, { forwardRef, useImperativeHandle, useState } from "react";
import { Modal, Form, Input, InputNumber, Select } from "antd";
import { addRiskStrategyApi, updateRiskStrategyApi, type RiskStrategyModel } from "@/api/app/risk/strategy";
import { message } from "@/hooks/useMessage";

export interface FormModalRef {
	showAddFormModal: () => void;
	showEditFormModal: (id: number, record?: RiskStrategyModel) => void;
}

interface Props {
	onSuccess?: () => void;
}

const statusOptions = [
	{ label: "启用", value: 1 },
	{ label: "停用", value: 0 }
];

const FormModal = forwardRef<FormModalRef, Props>(({ onSuccess }, ref) => {
	const [open, setOpen] = useState(false);
	const [confirmLoading, setConfirmLoading] = useState(false);
	const [isEdit, setIsEdit] = useState(false);
	const [currentId, setCurrentId] = useState<number>();
	const [form] = Form.useForm<RiskStrategyModel>();

	useImperativeHandle(ref, () => ({
		showAddFormModal: () => {
			setIsEdit(false);
			setCurrentId(undefined);
			form.resetFields();
			setOpen(true);
		},
		showEditFormModal: (id: number, record?: RiskStrategyModel) => {
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
				await updateRiskStrategyApi(currentId, values);
				message.success("更新成功");
			} else {
				await addRiskStrategyApi(values);
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
			title={isEdit ? "编辑策略" : "新增策略"}
			open={open}
			onOk={handleOk}
			onCancel={() => setOpen(false)}
			confirmLoading={confirmLoading}
			destroyOnClose
			width={600}
		>
			<Form form={form} layout="vertical">
				<Form.Item label="场景" name="scene" rules={[{ required: true, message: "请输入场景" }]}>
					<Input />
				</Form.Item>
				<Form.Item label="规则编码" name="ruleCode" rules={[{ required: true, message: "请输入规则编码" }]}>
					<Input />
				</Form.Item>
				<Form.Item label="统计维度" name="identityType" rules={[{ required: true, message: "请输入统计维度" }]}>
					<Input />
				</Form.Item>
				<Form.Item label="窗口秒" name="windowSeconds" rules={[{ required: true, message: "请输入窗口秒" }]}>
					<InputNumber min={1} style={{ width: "100%" }} />
				</Form.Item>
				<Form.Item label="阈值次数" name="threshold" rules={[{ required: true, message: "请输入阈值次数" }]}>
					<InputNumber min={1} style={{ width: "100%" }} />
				</Form.Item>
				<Form.Item label="动作编码" name="action" rules={[{ required: true, message: "请输入动作编码" }]}>
					<Input />
				</Form.Item>
				<Form.Item label="动作值" name="actionValue" rules={[{ required: true, message: "请输入动作值" }]}>
					<InputNumber min={0} style={{ width: "100%" }} />
				</Form.Item>
				<Form.Item label="状态" name="status" rules={[{ required: true, message: "请选择状态" }]}>
					<Select options={statusOptions} />
				</Form.Item>
				<Form.Item label="优先级" name="priority">
					<InputNumber min={0} style={{ width: "100%" }} />
				</Form.Item>
				<Form.Item label="说明" name="remark">
					<Input.TextArea rows={3} />
				</Form.Item>
			</Form>
		</Modal>
	);
});

FormModal.displayName = "FormModal";

export default FormModal;
