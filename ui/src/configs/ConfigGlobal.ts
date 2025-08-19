const StoreAuthLocation = {
  LocalStorage: 'localStorage',
  Cookie: 'cookie',
} as const;
type StoreAuthLocationType = (typeof StoreAuthLocation)[keyof typeof StoreAuthLocation];
declare global {
  interface ApisInternalModel {
    baseUrl: string;
  }

  interface Window {
    _apis: ApisInternalModel;
    _storeAuth: StoreAuthLocationType;
    _cdnPath: string;
    _basePath: string;
  }
}

export default {
  StoreAuthLocation,
} as const;
