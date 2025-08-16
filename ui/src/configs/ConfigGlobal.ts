declare global {
  interface ApisInternalModel {
    baseUrl: string;
  }

  interface Window {
    _apis: ApisInternalModel;
  }
}
