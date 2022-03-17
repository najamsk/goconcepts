// Code generated by command: go run asm.go -out tsip_amd64.s. DO NOT EDIT.

#include "textflag.h"

// func HashASM(k0 uint64, k1 uint64, p []byte) uint64
TEXT ·HashASM(SB), NOSPLIT, $0-48
	MOVQ k0+0(FP), AX
	MOVQ k1+8(FP), CX
	MOVQ $0x736f6d6570736575, DX
	XORQ DX, AX
	MOVQ $0x646f72616e646f6d, DX
	XORQ DX, CX
	MOVQ p_base+16(FP), DX
	MOVQ p_len+24(FP), BX
	MOVQ BX, BP
	SHLQ $0x38, BP
	CMPQ BX, $0x08
	JL   loop_end

loop_begin:
	MOVQ (DX), SI
	XORQ SI, CX
	ADDQ CX, AX
	ROLQ $0x0d, CX
	XORQ AX, CX
	ROLQ $0x23, AX
	ADDQ CX, AX
	ROLQ $0x11, CX
	XORQ AX, CX
	ROLQ $0x15, AX
	XORQ SI, AX
	ADDQ $0x08, DX
	SUBQ $0x08, BX
	CMPQ BX, $0x08
	JGE  loop_begin

loop_end:
	CMPQ BX, $0x00
	JE   sw0
	CMPQ BX, $0x01
	JE   sw1
	CMPQ BX, $0x02
	JE   sw2
	CMPQ BX, $0x03
	JE   sw3
	CMPQ BX, $0x04
	JE   sw4
	CMPQ BX, $0x05
	JE   sw5
	CMPQ BX, $0x06
	JE   sw6

sw7:
	MOVBQZX 6(DX), BX
	SHLQ    $0x30, BX
	ORQ     BX, BP

sw6:
	MOVBQZX 5(DX), BX
	SHLQ    $0x28, BX
	ORQ     BX, BP

sw5:
	MOVBQZX 4(DX), BX
	SHLQ    $0x20, BX
	ORQ     BX, BP

sw4:
	MOVBQZX 3(DX), BX
	SHLQ    $0x18, BX
	ORQ     BX, BP

sw3:
	MOVBQZX 2(DX), BX
	SHLQ    $0x10, BX
	ORQ     BX, BP

sw2:
	MOVBQZX 1(DX), BX
	SHLQ    $0x08, BX
	ORQ     BX, BP

sw1:
	MOVBQZX (DX), BX
	SHLQ    $0x00, BX
	ORQ     BX, BP

sw0:
	XORQ BP, CX
	ADDQ CX, AX
	ROLQ $0x0d, CX
	XORQ AX, CX
	ROLQ $0x23, AX
	ADDQ CX, AX
	ROLQ $0x11, CX
	XORQ AX, CX
	ROLQ $0x15, AX
	XORQ BP, AX
	XORQ $0xff, CX
	ADDQ CX, AX
	ROLQ $0x0d, CX
	XORQ AX, CX
	ROLQ $0x23, AX
	ADDQ CX, AX
	ROLQ $0x11, CX
	XORQ AX, CX
	ROLQ $0x15, AX
	ROLQ $0x20, CX
	ADDQ CX, AX
	ROLQ $0x0d, CX
	XORQ AX, CX
	ROLQ $0x23, AX
	ADDQ CX, AX
	ROLQ $0x11, CX
	XORQ AX, CX
	ROLQ $0x15, AX
	ROLQ $0x20, CX
	XORQ CX, AX
	MOVQ AX, ret+40(FP)
	RET
