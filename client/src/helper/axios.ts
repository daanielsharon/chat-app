import { API_URL } from "./../constant/index";
import axios from "axios";

const httpRequest = axios.create({
  baseURL: API_URL,
  timeout: 2000,
  timeoutErrorMessage: "Err: connection timeout",
  headers: {
    Accept: "application/json",
    "Content-Type": "application/json",
  },
});

export default httpRequest;
