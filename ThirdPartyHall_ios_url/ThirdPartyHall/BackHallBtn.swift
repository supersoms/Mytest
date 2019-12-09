//
//  IMOSuspendedBallView.swift
//  ThirdPartyHall
//
//  Created by LiuYun on 2019/3/28.
//  Copyright © 2019年 tianzi. All rights reserved.
//

import UIKit

private let ScreenHeight = UIScreen.main.bounds.size.height
private let ScreenWidth = UIScreen.main.bounds.size.width
private let cornerRadio:CGFloat = 30.0
private let placeWidth = 5.0
private let centerX:CGFloat = 30.0  //x半径
private let centerY:CGFloat = 30.0  //y半径

class BackHallBtn: UIView {
    static let ButtonSize: CGFloat = 60
    static let SubButtonsViewSize: CGFloat = 300
    var appWindow = UIWindow()//window
    let markView = UIButton(type: UIButton.ButtonType.custom)//黑色遮罩
    var superFrame = CGRect()//上级view的frame
    var hasSelected = false//是否展开
    var subButtonsView = UIButton(type: UIButton.ButtonType.custom)//展开的view
    var subButtonSize = CGSize()//展开页面里的button尺寸
    var superCon = ViewController()
//    var viewController : ViewController
    init(frame: CGRect, viewFrame: CGRect,con : ViewController) {
        var f = frame
        f.size.width = BackHallBtn.ButtonSize
        f.size.height = BackHallBtn.ButtonSize
        f.origin.x = frame.origin.x > viewFrame.size.width-frame.origin.x ? viewFrame.size.width-10-BackHallBtn.ButtonSize : 10
        super.init(frame: f)
        
        superFrame = viewFrame
        superCon = con
        setupButton()
        self.backgroundColor = UIColor.clear
        
        let panGesture = UIPanGestureRecognizer(target: self, action: #selector(panGestureAction(pan:)))
        self.isUserInteractionEnabled = true
        
        let tagGesture = UITapGestureRecognizer(target: self, action: #selector(showSubButtons))
        tagGesture.numberOfTapsRequired = 1
        
        panGesture.require(toFail: tagGesture)
        tagGesture.require(toFail: panGesture)
        
        self.addGestureRecognizer(tagGesture)
        self.addGestureRecognizer(panGesture)
        
        let appDelegate = (UIApplication.shared.delegate) as! AppDelegate
        appWindow = appDelegate.window!
        
        markView.frame = CGRect(x: 0, y: 0, width: appWindow.frame.size.width, height: appWindow.frame.size.height)
        markView.alpha = 0
        markView.backgroundColor = UIColor(white: 0, alpha: 0)
        markView.addTarget(self, action: #selector(showSelf), for: UIControl.Event.touchUpInside)
        
        subButtonsView.frame = CGRect(x: 0, y: 0, width: BackHallBtn.SubButtonsViewSize, height: BackHallBtn.SubButtonsViewSize)
        subButtonsView.layer.cornerRadius = 30
        subButtonsView.layer.backgroundColor = UIColor(white: 0, alpha: 0.85).cgColor
        subButtonsView.addTarget(self, action: #selector(showSelf), for: UIControl.Event.touchUpInside)
        hide()
    }
    
    required init?(coder aDecoder: NSCoder) {
        fatalError("init(coder:) has not been implemented")
    }
    
    func setupLaye(path: UIBezierPath, white: Float, alpha: Float) {
        let layer = CAShapeLayer()
        layer.path = path.cgPath
        layer.fillColor = UIColor(white: CGFloat(white), alpha: CGFloat(alpha)).cgColor
        layer.strokeColor = UIColor(white: CGFloat(white), alpha: CGFloat(alpha)).cgColor
        self.layer.addSublayer(layer)
    }
    
    func hide(){
        self.isHidden = true
    }
    
    func show(){
        self.isHidden = false
    }
    
    func setupButton() {
        let centerP = CGPoint(x:frame.size.width/2, y:frame.size.height/2)
        var path = UIBezierPath()
        path.addArc(withCenter: centerP, radius: BackHallBtn.ButtonSize/2, startAngle: 0, endAngle: CGFloat.pi*2, clockwise: true)
        setupLaye(path: path, white: 0, alpha: 0.85)
        
        path = UIBezierPath()
        path.addArc(withCenter: centerP, radius: BackHallBtn.ButtonSize/2.7, startAngle: 0, endAngle: CGFloat.pi*2, clockwise: true)
        setupLaye(path: path, white: 1, alpha: 0.2)
        
        path = UIBezierPath()
        path.addArc(withCenter: centerP, radius: BackHallBtn.ButtonSize/3.4, startAngle: 0, endAngle: CGFloat.pi*2, clockwise: true)
        setupLaye(path: path, white: 1, alpha: 0.3)
        
        path = UIBezierPath()
        path.addArc(withCenter: centerP, radius: BackHallBtn.ButtonSize/4.8, startAngle: 0, endAngle: CGFloat.pi*2, clockwise: true)
        setupLaye(path: path, white: 1, alpha: 1)
    }
    
    //拖拽
    @objc func panGestureAction(pan: UIPanGestureRecognizer) {
        if hasSelected {
            return
        }
        switch pan.state {
        case UIGestureRecognizer.State.began:
            self.superview?.bringSubviewToFront(self)
            break
        case UIGestureRecognizer.State.changed:
            let point = pan.translation(in: self)
            let f = self.frame
            let dx = point.x + f.origin.x
            let dy = point.y + f.origin.y
            
            self.frame = CGRect(x: dx, y: dy, width: BackHallBtn.ButtonSize, height: BackHallBtn.ButtonSize)
            //  注意一旦你完成上述的移动，将translation重置为0十分重要。否则translation每次都会叠加
            pan.setTranslation(CGPoint(x: 0, y: 0), in: self)
            break
        case UIGestureRecognizer.State.ended:
            let f = self.frame
            var dx = f.origin.x
            var dy = f.origin.y
            if dx > superFrame.size.width-10-BackHallBtn.ButtonSize {
                dx = superFrame.size.width-10-BackHallBtn.ButtonSize
            } else if dx < 10 {
                dx = 10.0
            } else {
                dx = dx > superFrame.size.width-dx ? superFrame.size.width-10-BackHallBtn.ButtonSize : 10
            }
            
            if dy > superFrame.size.height-10-BackHallBtn.ButtonSize {
                dy = superFrame.size.height-10-BackHallBtn.ButtonSize
            } else if dy < 10 {
                dy = 10.0
            }
            
            UIView.animate(withDuration: 0.2) {
                self.frame = CGRect(x: dx, y: dy, width: f.size.width, height: f.size.height)
            }
            break
            
        default: break
            
        }
        let point:CGPoint = pan.translation(in: self)
        self.transform = CGAffineTransform.init(translationX: point.x, y: point.y)
    }
    
    //点击
    @objc func showSubButtons() {
        superCon.backToHall()
    }
    
    @objc func showSelf() {
        UIView.animate(withDuration: 0.3, animations: {
            self.markView.alpha = 0
            self.subButtonsView.frame = CGRect(x: self.frame.origin.x, y: self.frame.origin.y, width: BackHallBtn.ButtonSize, height: BackHallBtn.ButtonSize)
        }) { (true) in
            self.isHidden = false
            self.hasSelected = false
            self.markView.removeFromSuperview()
            self.subButtonsView.removeFromSuperview()
        }
    }
}
