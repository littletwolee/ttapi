;;;;进来加载目录
(setq default-directory "~/code")


;;;;颜色配置  
(add-to-list 'custom-theme-load-path "~/.emacs.d/themes/")
(load-theme 'molokai t)
;(load-theme 'zenburn t)


;;;;行号
(require 'linum)
(global-linum-mode 1)


;;;;最大化  
;(defun my-maximized ()    
;  (interactive)    
;  (x-send-client-message    
;   nil 0 nil "_NET_WM_STATE" 32    
;   '(2 "_NET_WM_STATE_MAXIMIZED_HORZ" 0))    
;  (x-send-client-message    
;   nil 0 nil "_NET_WM_STATE" 32    
;   '(2 "_NET_WM_STATE_MAXIMIZED_VERT" 0)))    
;(my-maximized)  


(global-font-lock-mode t);语法高亮 
(setq visible-bell t) ;关闭错误提示音
(setq inhibit-startup-message t);关闭起动时LOGO
(setq default-major-mode 'erlang-mode);一打开就起用 text 模式 
(show-paren-mode t);显示括号匹配
(setq show-paren-style 'parenthesis);括号匹配时可以高亮显示另外一边的括号，但光标不会烦人的跳到另一个括号处。
(mouse-avoidance-mode 'animate) ;光标靠近鼠标指针时，让鼠标指针自动让开，别挡住视线。
(set-scroll-bar-mode nil);;取消滚动栏 
(auto-image-file-mode t);打开图片显示功能 
(fset 'yes-or-no-p 'y-or-n-p);以 y/n代表 yes/no 
(setq x-select-enable-clipboard t);支持emacs和外部程序的粘贴
(setq auto-save-default nil);关闭备份文件#xxx#
(setq-default cursor-type 'bar) ;设定光标为短线


;;;;全选
(defun select-all ()
  "Select the whole buffer."
  (interactive)
  (goto-char (point-min))
  ;; Mark current position and push it into the mark ring.
  (push-mark-command nil nil)
  (goto-char (point-max))
  (message "ok."))
(provide 'select-all)
(autoload 'select-all "select-all"
  "Select the whole buffer." t nil)
;; user defined keys
(global-set-key "\C-x\C-a" 'select-all)


;;;;auto-complete
(add-to-list 'load-path "~/.emacs.d/auto-complete")
(require 'auto-complete)
(require 'auto-complete-config)
(add-to-list 'ac-dictionary-directories "~/.emacs.d/auto-complete/ac-dict")
(ac-config-default)


;;;;erlang mode
(setq load-path (cons  "~/programs/otp/lib/erlang/lib/tools-2.8.2/emacs" load-path))
(setq erlang-root-dir "~/programs/otp/")
(setq exec-path (cons "~/programs/otp/bin" exec-path))
(require 'erlang-start)

 ;; Some Erlang customizations  
(add-hook 'erlang-mode-hook  
  (lambda ()  
  ;; when starting an Erlang shell in Emacs, default in the node name  
    (setq inferior-erlang-machine-options '("-sname" "emacs"))  
    ;; add Erlang functions to an imenu menu  
    (imenu-add-to-menubar "imenu")))  
;; A number of the erlang-extended-mode key bindings are useful in the shell too  
(defconst distel-shell-keys  
  '(("/C-/M-i"   erl-complete)  
    ("/M-?"      erl-complete)   
    ("/M-."      erl-find-source-under-point)  
    ("/M-,"      erl-find-source-unwind)   
    ("/M-*"      erl-find-source-unwind)   
    )  
  "Additional keys to bind when in Erlang shell.")  
(add-hook 'erlang-shell-mode-hook  
   (lambda ()  
     ;; add some Distel bindings to the Erlang shell  
     (dolist (spec distel-shell-keys)  
       (define-key erlang-shell-mode-map (car spec) (cadr spec)))))  


;;;;go
(add-to-list 'load-path "~/.emacs.d")
(require 'go-autocomplete)
(require 'auto-complete-config)
(require 'go-mode)


;;;;distel
(let ((distel-dir "~/.emacs.d/distel/elisp")) 
(unless (member distel-dir load-path) 
(setq load-path (append load-path (list distel-dir))))) 
(require 'distel) 
(distel-setup)


;;;;配置CEDET  
(load-file "~/.emacs.d/cedet/common/cedet.el")    ;你的安装目录
(global-ede-mode 1)
(semantic-load-enable-gaudy-code-helpers)
(global-srecode-minor-mode 1)


;;;;ecb
(add-to-list 'load-path "~/.emacs.d/ecb")    ;你的ecb解压目录
(require 'ecb)
(require 'ecb-autoloads)
(setq ecb-auto-activate t)
(setq ecb-tip-of-the-day nil)
(global-set-key (kbd "C-<f7>") 'ecb-minor-mode)   ; 打开ejb
;ejb 快捷键
(global-set-key (kbd "C-<left>") 'windmove-left)   ;左边窗口
(global-set-key (kbd "C-<right>") 'windmove-right)  ;右边窗口
(global-set-key (kbd "C-<up>") 'windmove-up)     ; 上边窗口
(global-set-key (kbd "C-<down>") 'windmove-down)   ; 下边窗口
(setq stack-trace-on-error t) ;”Symbol's value as variable is void: stack-trace-on-error“
;ceb鼠标点选
(custom-set-variables '(ecb-primary-secondary-mouse-buttons (quote mouse-1--mouse-2))
 '(ecb-source-path (quote (("/" "/")))))
(custom-set-faces)


;;;;yasnippet
 (add-to-list 'load-path "~/.emacs.d/yasnippet")
    (require 'yasnippet) ;; not yasnippet-bundle
;    (yas/initialize)
    (yas/load-directory "~/.emacs.d/yasnippet/snippets")
(setq yas/prompt-functions 
   '(yas/dropdown-prompt yas/x-prompt yas/completing-prompt yas/ido-prompt yas/no-prompt))
(yas/global-mode 1)
(yas/minor-mode-on) ; 以minor mode打开，这样才能配合主mode使用
