import React, { forwardRef, useImperativeHandle, useState } from "react";
import { Modal, Form, Input, InputNumber, Select } from "antd";
import { addConfigApi, updateConfigApi, type ChallengeConfigModel } from "@/api/app/challenge/config";
import { message } from "@/hooks/useMessage";

export interface FormModalRef {
  showAddFormModal: () => void;
  showEditFormModal: (id: number, record?: ChallengeConfigModel) => void;
}

interface Props {
  onSuccess?: () => void;
}

const statusOptions = [
  { label: "启用", value: 1 },
  { label: "禁用", value: 0 }
];

const FormModal = forwardRef<FormModalRef, Props>(({ onSuccess }, ref) => {
  const [open, setOpen] = useState(false);
  const [confirmLoading, setConfirmLoading] = useState(false);
  const [isEdit, setIsEdit] = useState(false);
  const [currentId, setCurrentId] = useState<number>();
  const [form] = Form.useForm<ChallengeConfigModel>();

  useImperativeHandle(ref, () => ({
    showAddFormModal: () => {
      setIsEdit(false);
      setCurrentId(undefined);
      form.resetFields();
      setOpen(true);
    },
    showEditFormModal: (id: number, record?: ChallengeConfigModel) => {
      setIsEdit(true);
      setCurrentId(id);
      if (record) {
        form.setFieldsValue(record);
      }
      setOpen(true);
    }
  }));

  const handleOk = async () => {
    try {
      const values = await form.validateFields();
      setConfirmLoading(true);
      if (isEdit && currentId) {
        await updateConfigApi(currentId, values);
        message.success("更新成功");
      } else {
        await addConfigApi(values);
        message.success("新增成功");
      }
      setOpen(false);
      onSuccess?.();
    } catch (e) {
      // ignore validation errors
    } finally {
      setConfirmLoading(false);
    }
  };

  return (
    <Modal
      title={isEdit ? "编辑挑战配置" : "新增挑战配置"}
      open={open}
      onOk={handleOk}
      onCancel={() => setOpen(false)}
      confirmLoading={confirmLoading}
      destroyOnClose
    >
      <Form form={form} layout="vertical">
        <Form.Item label="挑战天数" name="dayCount" rules={[{ required: true, message: "请输入挑战天数" }]}>
          <InputNumber min={1} style={{ width: "100%" }} />
        </Form.Item>
        <Form.Item label="挑战金额" name="amount" rules={[{ required: true, message: "请输入挑战金额" }]}>
          <Input />
        </Form.Item>
        <Form.Item label="打卡开始时间" name="checkinStart" rules={[{ required: true, message: "请输入打卡开始时间" }]}>
          <Input placeholder="例如 08:00:00" />
        </Form.Item>
        <Form.Item label="打卡结束时间" name="checkinEnd" rules={[{ required: true, message: "请输入打卡结束时间" }]}>
          <Input placeholder="例如 23:00:00" />
        </Form.Item>
        <Form.Item label="平台奖励" name="platformBonus">
          <Input />
        </Form.Item>
        <Form.Item label="状态" name="status" rules={[{ required: true, message: "请选择状态" }]}>
          <Select options={statusOptions} />
        </Form.Item>
        <Form.Item label="排序" name="sort">
          <InputNumber min={0} style={{ width: "100%" }} />
        </Form.Item>
      </Form>
    </Modal>
  );
});

FormModal.displayName = "FormModal";

export default FormModal;
