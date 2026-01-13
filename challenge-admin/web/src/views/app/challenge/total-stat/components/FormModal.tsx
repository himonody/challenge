import React, { forwardRef, useImperativeHandle, useState } from "react";
import { Modal, Form, Input, InputNumber } from "antd";
import { addTotalStatApi, updateTotalStatApi, type ChallengeTotalStatModel } from "@/api/app/challenge/totalStat";
import { message } from "@/hooks/useMessage";

export interface FormModalRef {
	showAddFormModal: () => void;
	showEditFormModal: (id: number, record?: ChallengeTotalStatModel) => void;
}

interface Props {
	onSuccess?: () => void;
}

const FormModal = forwardRef<FormModalRef, Props>(({ onSuccess }, ref) => {
	const [open, setOpen] = useState(false);
	const [confirmLoading, setConfirmLoading] = useState(false);
	const [isEdit, setIsEdit] = useState(false);
	const [currentId, setCurrentId] = useState<number>();
	const [form] = Form.useForm<ChallengeTotalStatModel>();

	useImperativeHandle(ref, () => ({
		showAddFormModal: () => {
			setIsEdit(false);
			setCurrentId(undefined);
			form.resetFields();
			setOpen(true);
		},
		showEditFormModal: (id: number, record?: ChallengeTotalStatModel) => {
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
				await updateTotalStatApi(currentId, values);
				message.success("更新成功");
			} else {
				await addTotalStatApi(values);
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
			title={isEdit ? "编辑总统计" : "新增总统计"}
			open={open}
			onOk={handleOk}
			onCancel={() => setOpen(false)}
			confirmLoading={confirmLoading}
			destroyOnClose
		>
			<Form form={form} layout="vertical">
				<Form.Item label="ID" name="id" rules={[{ required: true, message: "请输入ID" }]}>
					<InputNumber min={0} style={{ width: "100%" }} disabled={isEdit} />
				</Form.Item>
				<Form.Item label="总用户数" name="totalUserCnt">
					<InputNumber min={0} style={{ width: "100%" }} />
				</Form.Item>
				<Form.Item label="总参与次数" name="totalJoinCnt">
					<InputNumber min={0} style={{ width: "100%" }} />
				</Form.Item>
				<Form.Item label="总成功次数" name="totalSuccessCnt">
					<InputNumber min={0} style={{ width: "100%" }} />
				</Form.Item>
				<Form.Item label="总失败次数" name="totalFailCnt">
					<InputNumber min={0} style={{ width: "100%" }} />
				</Form.Item>
				<Form.Item label="总参与金额" name="totalJoinAmount">
					<Input />
				</Form.Item>
				<Form.Item label="总成功金额" name="totalSuccessAmount">
					<Input />
				</Form.Item>
				<Form.Item label="总失败金额" name="totalFailAmount">
					<Input />
				</Form.Item>
				<Form.Item label="总平台奖励" name="totalPlatformBonus">
					<Input />
				</Form.Item>
				<Form.Item label="总奖池金额" name="totalPoolAmount">
					<Input />
				</Form.Item>
			</Form>
		</Modal>
	);
});

FormModal.displayName = "FormModal";

export default FormModal;
