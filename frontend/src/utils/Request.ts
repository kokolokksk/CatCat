
import axios, { AxiosRequestConfig, AxiosResponse } from 'axios';

export const makeRequest = async <T>(config: AxiosRequestConfig): Promise<AxiosResponse<T>> => {
   
    const response = await axios(config);
    return response.data;
  
};
