import { ListCondition } from "../pb/common/request.ts";
import { LoginResponse } from "../pb/user/response.ts";
import { TableProps } from "antd";
export interface LoginProps {
  onSuccess: (resp: LoginResponse) => void;
  onLoading: () => void;
  onFailure: (message: string) => void;
}

export interface ListResponse<Record> {
  items: Array<Record>;
  total: number;
}

export interface STableProps<Record = any> extends TableProps<Record>{
  data: (listCondition: ListCondition) => Promise<ListResponse<Record>>
}