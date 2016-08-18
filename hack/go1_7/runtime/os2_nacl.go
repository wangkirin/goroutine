// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

const (
	_NSIG		= 32
	_SI_USER	= 1

	_EPERM		= 1
	_ENOENT		= 2
	_ESRCH		= 3
	_EINTR		= 4
	_EIO		= 5
	_ENXIO		= 6
	_E2BIG		= 7
	_ENOEXEC	= 8
	_EBADF		= 9
	_ECHILD		= 10
	_EAGAIN		= 11

	_EACCES			= 13
	_EFAULT			= 14
	_EBUSY			= 16
	_EEXIST			= 17
	_EXDEV			= 18
	_ENODEV			= 19
	_ENOTDIR		= 20
	_EISDIR			= 21
	_EINVAL			= 22
	_ENFILE			= 23
	_EMFILE			= 24
	_ENOTTY			= 25
	_EFBIG			= 27
	_ENOSPC			= 28
	_ESPIPE			= 29
	_EROFS			= 30
	_EMLINK			= 31
	_EPIPE			= 32
	_ENAMETOOLONG		= 36
	_ENOSYS			= 38
	_EDQUOT			= 122
	_EDOM			= 33
	_ERANGE			= 34
	_EDEADLK		= 35
	_ENOLCK			= 37
	_ENOTEMPTY		= 39
	_ELOOP			= 40
	_ENOMSG			= 42
	_EIDRM			= 43
	_ECHRNG			= 44
	_EL2NSYNC		= 45
	_EL3HLT			= 46
	_EL3RST			= 47
	_ELNRNG			= 48
	_EUNATCH		= 49
	_ENOCSI			= 50
	_EL2HLT			= 51
	_EBADE			= 52
	_EBADR			= 53
	_EXFULL			= 54
	_ENOANO			= 55
	_EBADRQC		= 56
	_EBADSLT		= 57
	_EDEADLOCK		= _EDEADLK
	_EBFONT			= 59
	_ENOSTR			= 60
	_ENODATA		= 61
	_ETIME			= 62
	_ENOSR			= 63
	_ENONET			= 64
	_ENOPKG			= 65
	_EREMOTE		= 66
	_ENOLINK		= 67
	_EADV			= 68
	_ESRMNT			= 69
	_ECOMM			= 70
	_EPROTO			= 71
	_EMULTIHOP		= 72
	_EDOTDOT		= 73
	_EBADMSG		= 74
	_EOVERFLOW		= 75
	_ENOTUNIQ		= 76
	_EBADFD			= 77
	_EREMCHG		= 78
	_ELIBACC		= 79
	_ELIBBAD		= 80
	_ELIBSCN		= 81
	_ELIBMAX		= 82
	_ELIBEXEC		= 83
	_EILSEQ			= 84
	_EUSERS			= 87
	_ENOTSOCK		= 88
	_EDESTADDRREQ		= 89
	_EMSGSIZE		= 90
	_EPROTOTYPE		= 91
	_ENOPROTOOPT		= 92
	_EPROTONOSUPPORT	= 93
	_ESOCKTNOSUPPORT	= 94
	_EOPNOTSUPP		= 95
	_EPFNOSUPPORT		= 96
	_EAFNOSUPPORT		= 97
	_EADDRINUSE		= 98
	_EADDRNOTAVAIL		= 99
	_ENETDOWN		= 100
	_ENETUNREACH		= 101
	_ENETRESET		= 102
	_ECONNABORTED		= 103
	_ECONNRESET		= 104
	_ENOBUFS		= 105
	_EISCONN		= 106
	_ENOTCONN		= 107
	_ESHUTDOWN		= 108
	_ETOOMANYREFS		= 109
	_ETIMEDOUT		= 110
	_ECONNREFUSED		= 111
	_EHOSTDOWN		= 112
	_EHOSTUNREACH		= 113
	_EALREADY		= 114
	_EINPROGRESS		= 115
	_ESTALE			= 116
	_ENOTSUP		= _EOPNOTSUPP
	_ENOMEDIUM		= 123
	_ECANCELED		= 125
	_ELBIN			= 2048
	_EFTYPE			= 2049
	_ENMFILE		= 2050
	_EPROCLIM		= 2051
	_ENOSHARE		= 2052
	_ECASECLASH		= 2053
	_EWOULDBLOCK		= _EAGAIN

	_PROT_NONE	= 0x0
	_PROT_READ	= 0x1
	_PROT_WRITE	= 0x2
	_PROT_EXEC	= 0x4

	_MAP_SHARED	= 0x1
	_MAP_PRIVATE	= 0x2
	_MAP_FIXED	= 0x10
	_MAP_ANON	= 0x20

	_MADV_FREE	= 0
	_SIGFPE		= 8
	_FPE_INTDIV	= 0
)

type siginfo struct{}
