/**
 * Namespace Api
 *
 * All backend api type
 */
declare namespace Api {
  namespace Common {
    /** common params of paginating */
    interface PaginatingCommonParams {
      /** current page number */
      current: number;
      /** page size */
      size: number;
      /** total count */
      total: number;
    }

    /** common params of paginating query list data */
    interface PaginatingQueryRecord<T = any> extends PaginatingCommonParams {
      records: T[];
    }

    /** common search params of table */
    type CommonSearchParams = Pick<Common.PaginatingCommonParams, 'current' | 'size'>;

    /** common record */
    type CommonRecord<T = any> = {
      id?: number;
      created_at?: string;
    } & T;
  }

  namespace Form {
    interface LoginForm {
      username: string;
      password: string;
      code: string;
      captchaId: string;
    }
  }

  /**
   * namespace Auth
   *
   * backend api module: "auth"
   */
  namespace Auth {
    interface LoginToken {
      token: string;
    }

    interface Captcha {
      id: string;
      img: string;
    }

    interface UserInfo {
      userId: string;
      userName: string;
      roles: string[];
      buttons: string[];
    }
  }

  namespace Home {
    interface Stat {
      userCount: number;
      peerCount: number;
      onlineCount: number;
      visitsCount: number;
    }

    interface LineChart {
      xAxis: string[];
      users: number[];
      peer: number[];
    }

    interface PieChart {
      name: string;
      value: number;
    }
  }

  namespace UserManagement {
    type User = Common.CommonRecord<{
      username: string;
      password: string;
      name: string;
      email: string;
      licensed_devices: number;
      note: string;
      status: number;
      is_admin: boolean;
      admin_status: number;
    }>;
    type UserList = Common.PaginatingQueryRecord<User>;

    type UserSearchParams = CommonType.RecordNullable<
      Pick<Api.UserManagement.User, 'username' | 'name' | 'email' | 'status' | 'admin_status' | 'created_at'> &
        Api.Common.CommonSearchParams
    >;

    type Session = Common.CommonRecord<{
      username: string;
      rustdesk_id: string;
      expired: string;
    }>;
    type SessionList = Common.PaginatingQueryRecord<Session>;

    type SessionSearchParams = CommonType.RecordNullable<
      Pick<Api.UserManagement.Session, 'username' | 'created_at'> & Api.Common.CommonSearchParams
    >;
  }

  /**
   * namespace Route
   *
   * backend api module: "route"
   */
  namespace Route {
    type ElegantConstRoute = import('@elegant-router/types').ElegantConstRoute;

    interface MenuRoute extends ElegantConstRoute {
      id: string;
    }

    interface UserRoute {
      routes: MenuRoute[];
      home: import('@elegant-router/types').LastLevelRouteKey;
    }
  }
}
