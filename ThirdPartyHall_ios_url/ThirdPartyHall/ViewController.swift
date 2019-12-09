//  Created by LiuYun on 2019/3/27.
//  Copyright © 2019年 tianzi. All rights reserved.

import UIKit
import WebKit

class ViewController: UIViewController,WKUIDelegate,WKScriptMessageHandler,WKNavigationDelegate {
    
    func userContentController(_ userContentController: WKUserContentController, didReceive message: WKScriptMessage) {
        switch message.name {
        case "copyText":
            print("copy text by js sending message")
            UIPasteboard.general.string = message.body as? String
            showAlertMessage("复制成功", showtime: 2)
            break
        case "saveImg":
            self.saveImgFun(imgData: message.body as! String)
            break
        case "showImageView":
            if backHallBtn != nil{
                backHallBtn.show()
            }
            break
        case "hideImageView":
            if backHallBtn != nil{
                backHallBtn.hide()
            }
            break
        case "toOnlineByUrl":
            self.toOnlineByUrl(url: message.body as! String)
            break
        case "printLog":
            print(message.body)
            break
        default:
            break
        }
    }
    
    //返回大厅 这里要调用js代码
    func backToHall(){
        webView.evaluateJavaScript("close_child()", completionHandler: nil)
    }
    
    //保存图片
    func saveImgFun(imgData : String){
        let base64String = imgData.replacingOccurrences(of: "data:image/png;base64,", with:"")
        var imageData = Data(base64Encoded: base64String,options: .ignoreUnknownCharacters)
        if imageData == nil{
            imageData = Data(base64Encoded: base64String + "==",options: .ignoreUnknownCharacters)
        }
        var image:UIImage?
        if imageData != nil{
            image = UIImage(data: imageData!)
        }
        if image != nil{
            UIImageWriteToSavedPhotosAlbum(image!, self, #selector(saveCallBack(image:didFinishSavingWithError:contextInfo:)), nil)
        }
    }
    
    //保存图片回调方法
    @objc func saveCallBack(image:UIImage, didFinishSavingWithError:NSError?,contextInfo:AnyObject) {
        if didFinishSavingWithError != nil {
            showAlertMessage("保存失败", showtime: 2)
        } else {
            showAlertMessage("保存成功", showtime: 2)
        }
    }

    //跳转到Safari去打开url网页
    func toOnlineByUrl(url: String){
        if let url = URL(string: url) {
            //根据iOS系统版本，分别处理
            if #available(iOS 10, *) {
                UIApplication.shared.open(url, options: [:],
                                          completionHandler: {
                                            (success) in
                })
            } else {
                UIApplication.shared.openURL(url)
            }
        }
    }
    var webView: WKWebView!

    override func loadView() {                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   
        let webConfiguration = WKWebViewConfiguration()
        webView = WKWebView(frame: .zero, configuration: webConfiguration)
        webView.uiDelegate = self
        webView.navigationDelegate = self
        view = webView
        view.frame.size.width = UIScreen.main.bounds.size.width;
        view.frame.size.height = UIScreen.main.bounds.size.height;
        webView.scrollView.contentInsetAdjustmentBehavior = .never
        webView.scrollView.isScrollEnabled = false
        webConfiguration.userContentController.add(self, name: "copyText")
        webConfiguration.userContentController.add(self, name: "saveImg")
        webConfiguration.userContentController.add(self, name: "showImageView")
        webConfiguration.userContentController.add(self, name: "hideImageView")
        webConfiguration.userContentController.add(self, name: "toOnlineByUrl")
        webConfiguration.userContentController.add(self, name: "printLog")
    }
    
    var backHallBtn : BackHallBtn!
    override func viewDidLoad() {
        super.viewDidLoad()
        let gameUrl = Bundle.main.infoForKey("GAME_URL")!
        let myURL = URL(string: gameUrl) //"http://game.qsgames.la/ThirdPartyHall?type=1"
        let myRequest = URLRequest(url: myURL!)
        webView.load(myRequest)
        
        //仿苹果的全局浮动按钮
        var frame: CGRect = view.frame
        frame.size.height = UIScreen.main.bounds.size.height
        frame.size.width = UIScreen.main.bounds.size.width
        if backHallBtn == nil{
            backHallBtn = BackHallBtn(frame: CGRect(x: 100, y: 100, width: 100, height: 100), viewFrame: frame,con: self)
            view.addSubview(backHallBtn)
        }
    }
    
    //webview 加载内容失败
    func webView(_ webView: WKWebView, didFailProvisionalNavigation navigation: WKNavigation!, withError error: Error) {
        let alert = UIAlertController(title: "提示", message: "网络连接失败，请重新启动",preferredStyle: UIAlertController.Style.alert)
        let btnOK = UIAlertAction(title: "好的", style: .default, handler: {
            action in
            exit(0)
        })
        alert.addAction(btnOK)
        self.present(alert, animated: true, completion: nil)
    }
}

//显示弹出信息
func showAlertMessage(_ str:String,showtime Num:Double){
    let alert = UIAlertView(title: str, message: nil, delegate: nil, cancelButtonTitle: nil);
    alert.show()
    //        self.performSelector(#selector(dismissAlert(_:)), withObject: alert, afterDelay: Num)
    let dispatchTime: DispatchTime = DispatchTime.now() + Double(Int64(0.10 * Num * Double(NSEC_PER_SEC))) / Double(NSEC_PER_SEC)
    DispatchQueue.main.asyncAfter(deadline: dispatchTime, execute: {
        dismissAlert(alert)
    })
}

func dismissAlert(_ alert:UIAlertView){
    alert.dismiss(withClickedButtonIndex: alert.cancelButtonIndex, animated: true)
}
