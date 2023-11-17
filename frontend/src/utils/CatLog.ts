import { Log } from "../../wailsjs/go/main/App";

class CatLog {
  public static info = (args: any) => {
    //window.electron.ipcRenderer.sendMessage('log', ['info', args]);
    Log(args,'info')
  };

  public static error = (args: any) => {
    //window.electron.ipcRenderer.sendMessage('log', ['error', args]);
    Log(args,'error')
  };

  public static warn = (args: any) => {
    //window.electron.ipcRenderer.sendMessage('log', ['warn', args]);
    Log(args,'error')
  };

  public static debug = (args: any) => {
    //window.electron.ipcRenderer.sendMessage('log', ['debug', args]);
    Log(args,'error')
  };

  public static log = (args: any) => {
   // window.electron.ipcRenderer.sendMessage('log', ['log', args]);
   Log(args,'error')
  };

  public static console = (args?: any, ...optionalParams: any[]) => {
    console.log(args, ...optionalParams);
  };
}

export default CatLog;
// # sourceMappingURL=CatLog.js.map
// CONCATENATED MODULE: ./src/renderer/utils/CatLog.ts
