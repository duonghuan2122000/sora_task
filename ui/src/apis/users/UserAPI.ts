import HttpBase, { RequestTimeoutError, type BaseRequest, type BaseResponse } from '../HttpBase';
import type { LoginUserRequest, LoginUserResponse } from './UserModel';

/**
 * Thực hiện gửi request đăng nhập
 */
const login = async (payload: LoginUserRequest): Promise<BaseResponse<LoginUserResponse>> => {
  try {
    let reqBody: BaseRequest<LoginUserRequest> = {
      data: {
        attributes: payload,
      },
    };
    const res = await HttpBase.makeRequest<BaseResponse<LoginUserResponse>>({
      method: 'POST',
      url: '/users/login/by-mail',
      data: reqBody,
      timeout: 5000,
    });
    return res;
  } catch (error) {
    console.log(error);
    if (error instanceof RequestTimeoutError) {
      return {
        status: false,
        error: {
          code: '408',
          message: 'Timeout khi thực hiện request',
        },
      };
    }
    return {
      status: false,
      error: {
        code: '500',
        message: 'Có lỗi xảy ra',
      },
    };
  }
};

export default {
  login,
} as const;
