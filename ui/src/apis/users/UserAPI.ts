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
    });
    return res;
  } catch (error) {
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

const verifyUser = async () => {
  try {
    const res = await HttpBase.makeRequest<BaseResponse<void>>({
      method: 'GET',
      url: '/users/verify',
    });
    return res;
  } catch (error) {
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
  verifyUser,
} as const;
