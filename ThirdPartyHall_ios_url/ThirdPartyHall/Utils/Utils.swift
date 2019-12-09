import Foundation

class Utils {
    
    static let shared = Utils()
    private let defaults = UserDefaults.standard  //获取UserDefaults实例
    private init() {
        
    }
}

extension Utils {
    func saveSetting(value: Int, key: String){
        defaults.set(value, forKey: key)
        defaults.synchronize()
    }
    
    func saveSetting(value: String, key: String){
        defaults.set(value, forKey: key)
        defaults.synchronize()
    }
    
    func saveSetting(value: Bool, key: String){
        defaults.set(value, forKey: key)
        defaults.synchronize()
    }
    
    func getInt(key: String) -> Int {
        return defaults.integer(forKey: key)
    }
    
    func getString(key: String) -> String {
        return defaults.string(forKey: key) ?? ""
    }
    
    func getBool(key: String) -> Bool {
        return defaults.bool(forKey : key)
    }
    
    func removeData(key: String) {
        defaults.removeObject(forKey: key)
    }
}
