import { keys } from "../constants/globals.constant";
import ShowNotification from './notification.util';

export const callApi = async (
    options,
    headers = { "Content-type": "application/json" },
    isCompleteURL = false
  ) => {
    try {
      const { endpoint, config } = options;
      config.headers = headers;
      const url = isCompleteURL ? endpoint : keys.hostName + endpoint;
      const response = await fetch(url, config)
      const jsonResponse = await response.json();
      if(!jsonResponse.status) {
        if(response.status === 401)
          ShowNotification("Invalid authentication details", "error")
        else if (response.status === 403)
          ShowNotification("You're not authorized to access this page.", "error")
        else if (response.status === 500) {
          if(jsonResponse.message)
            ShowNotification(jsonResponse.message, "error")
          else
            ShowNotification("Internal server error. Please try again later!", "error")
        }
      }
      return jsonResponse;
    } catch (err) {
  
    }
  };
  