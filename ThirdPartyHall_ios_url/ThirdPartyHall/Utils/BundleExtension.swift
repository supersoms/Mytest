import Foundation

extension Bundle {
    //根据key从 .xcconfig 配置文件中读取配置信息
    func infoForKey(_ key: String) -> String? {
        return (infoDictionary?[key] as? String)?.replacingOccurrences(of: "\\", with: "")
    }
}
