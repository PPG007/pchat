import { ValidateMessages } from "rc-field-form/lib/interface";
import i18n from "../../i18n";

export const validateMessage: ValidateMessages = {
  required: i18n.t('validateMessage.required'),
  types: {
    email: i18n.t('validateMessage.email')
  }
}