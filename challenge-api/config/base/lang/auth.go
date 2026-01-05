package lang

const (
	AuthUsernameErrorCode           = 10000 // 用户名格式错误
	AuthPasswordErrorCode           = 10001 // 密码格式错误
	AuthVerificationCodeErrorCode   = 10002 // 验证码错误
	AuthUserAlreadyExistsCode       = 10003 // 用户已存在
	AuthInviteCodeNotFoundErrorCode = 10004 // 推荐码不存在

	//SysDictDataValueEmptyCode     = 10005
	//SysDictDataSortEmptyCode      = 10006
	//SysDictDataValueExistCode     = 10007
	//
	//// 配置管理
	//SysConfNameEmptyCode       = 10100
	//SysConfKeyEmptyCode        = 10101
	//SysConfValueEmptyCode      = 10102
	//SysConfTypeEmptyCode       = 10103
	//SysConfIsFrontendEmptyCode = 10104
	//SysConfKeyExistCode        = 10105
	//SysConfGetErrLogCode       = 10106
	//SysConfGetErrCode          = 10107
	//
	////部门管理
	//SysDeptParentIdEmptyCode   = 10200
	//SysDeptNameEmptyCode       = 10201
	//SysDeptLeaderEmptyCode     = 10202
	//SysDeptNameExistCode       = 10203
	//SysDeptChildExistNoDelCode = 10204
	//SysDeptParentSelfCode      = 10205
	//
	////角色管理
	//SysRoleNameEmptyCode         = 10301
	//SysRoleStatusEmptyCode       = 10302
	//SysRoleKeyEmptyCode          = 10303
	//SysRoleSortEmptyCode         = 10304
	//SysRoleKeyExistCode          = 10305
	//SysRoleAdminNoOpCode         = 10306
	//SysRoleUserExistNoDeleteCode = 10307
	//
	////岗位管理
	//SysPostNameEmptyCode   = 10400
	//SysPostCodeEmptyCode   = 10401
	//SysPostSortEmptyCode   = 10402
	//SysPostStatusEmptyCode = 10403
	//SysPostNameExistCode   = 10404
	//
	////用户管理
	//SysUserNameEmptyCode              = 10500
	//SysNickNameEmptyCode              = 10501
	//SysUserPhoneEmptyCode             = 10502
	//SysUserEmailEmptyCode             = 10503
	//SysUserDeptEmptyCode              = 10504
	//SysUserPwdEmptyCode               = 10505
	//SysUserNameExistCode              = 10506
	//SysUserNickNameExistCode          = 10507
	//SysUserPhoneExistCode             = 10508
	//SysUserEmailExistCode             = 10509
	//SysUserEmailFormatErrCode         = 10510
	//SysUserStatusEmptyCode            = 10511
	//SysUserNewPwdEmptyCode            = 10512
	//SysUserPwdErrCode                 = 10513
	//SysUserNoExistCode                = 10514
	//SysUserAvatarErrCode              = 10515
	//SysUserAvatarErrLogCode           = 10516
	//SysUseCapErrLogCode               = 10517
	//SysUseLogoutSuccessCode           = 10518
	//SysUseLoginOpCode                 = 10519
	//SysUseLogoutOpCode                = 10520
	//SysUseGenCaptchaErrCode           = 10521
	//SysUseGenCaptchaErrLogCode        = 10522
	//SysUseAvatarUploadErrCode         = 10523
	//SysUseAvatarUploadErrLogCode      = 10524
	//SysAdminUserNotAllowDeleteErrCode = 10525
	//SysUserNoRoleErrCode              = 10526
	//
	////菜单
	//SysMenuParentIdEmptyCode = 10600
	//SysMenuTitleEmptyCode    = 10601
	//SysMenuTypeEmptyCode     = 10602
	//SysMenuSortEmptyCode     = 10603
	//SysMenuHasChildCode      = 10604
	//SysNoRoleMenuCode        = 10605
	//SysMenuPathExistCode     = 10606
	//
	////gen表
	//SysGenTableSelectCode               = 10700
	//SysGenTableInsertExistCode          = 10701
	//SysGenTableImportErrLogCode         = 10702
	//SysGenTableImportErrCode            = 10703
	//SysGenTemplateModelReadErrCode      = 10704
	//SysGenTemplateModelReadLogErrCode   = 10705
	//SysGenTemplateApiReadErrCode        = 10706
	//SysGenTemplateApiReadLogErrCode     = 10707
	//SysGenTemplateJsReadErrCode         = 10708
	//SysGenTemplateJsReadLogErrCode      = 10709
	//SysGenTemplateRouterReadErrCode     = 10712
	//SysGenTemplateRouterReadLogErrCode  = 10713
	//SysGenTemplateDtoReadErrCode        = 10714
	//SysGenTemplateDtoReadLogErrCode     = 10715
	//SysGenTemplateServiceReadErrCode    = 10716
	//SysGenTemplateServiceReadLogErrCode = 10717
	//SysGenCreatePathLogErrCode          = 10718
	//SysGenCreatePathErrCode             = 10719
	//SysGenTemplateModelDecodeErrCode    = 10720
	//SysGenTemplateModelDecodeLogErrCode = 10721
	//
	////API
	//SysApiDirGetLogErrCode = 10803
	//SysApiDirGetErrCode    = 10804
	//SysApiParseLogErrCode  = 10805
	//SysApiParseErrCode     = 10806
)
