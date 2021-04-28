import { message } from '../components/antd/antd';

export default function (content, type = "success", duration = 3) {
  switch (type) {
    case 'success':
      message.success(content, duration);
      break;
    case 'error':
      message.error(content, duration);
      break;
    case 'info':
      message.info(content, duration);
      break;
    case 'warning':
      message.warning(content, duration);
      break;
    default: break;
  }
}