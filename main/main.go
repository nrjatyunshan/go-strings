package main

import "log"

/*
# objdump -d main | awk -v RS= '/^[[:xdigit:]]+ <main.TestLogPrintf>/'
000000000048ade0 <main.TestLogPrintf>:
  48ade0:       49 3b 66 10             cmp    0x10(%r14),%rsp
  48ade4:       76 72                   jbe    48ae58 <main.TestLogPrintf+0x78>
  48ade6:       48 83 ec 50             sub    $0x50,%rsp
  48adea:       48 89 6c 24 48          mov    %rbp,0x48(%rsp)
  48adef:       48 8d 6c 24 48          lea    0x48(%rsp),%rbp
  48adf4:       44 0f 11 7c 24 28       movups %xmm15,0x28(%rsp)
  48adfa:       44 0f 11 7c 24 38       movups %xmm15,0x38(%rsp)
  48ae00:       48 8d 15 39 76 00 00    lea    0x7639(%rip),%rdx        # 492440 <type.*+0x7440>
  48ae07:       48 89 54 24 28          mov    %rdx,0x28(%rsp)
  48ae0c:       48 8d 15 e5 79 03 00    lea    0x379e5(%rip),%rdx        # 4c27f8 <runtime.defaultGOROOT.str+0x70>
  48ae13:       48 89 54 24 30          mov    %rdx,0x30(%rsp)
  48ae18:       48 8d 15 a1 6f 00 00    lea    0x6fa1(%rip),%rdx        # 491dc0 <type.*+0x6dc0>
  48ae1f:       48 89 54 24 38          mov    %rdx,0x38(%rsp)
  48ae24:       48 8d 15 05 79 03 00    lea    0x37905(%rip),%rdx        # 4c2730 <$f64.bfe62e42fefa39ef+0x10>
  48ae2b:       48 89 54 24 40          mov    %rdx,0x40(%rsp)
  48ae30:       48 8d 05 12 9d 01 00    lea    0x19d12(%rip),%rax        # 4a4b49 <go.string.*+0x29b1>
  48ae37:       bb 13 00 00 00          mov    $0x13,%ebx
  48ae3c:       48 8d 4c 24 28          lea    0x28(%rsp),%rcx
  48ae41:       bf 02 00 00 00          mov    $0x2,%edi
  48ae46:       48 89 fe                mov    %rdi,%rsi
  48ae49:       e8 92 fd ff ff          call   48abe0 <log.Printf>
  48ae4e:       48 8b 6c 24 48          mov    0x48(%rsp),%rbp
  48ae53:       48 83 c4 50             add    $0x50,%rsp
  48ae57:       c3                      ret
  48ae58:       e8 43 08 fd ff          call   45b6a0 <runtime.morestack_noctxt.abi0>
  48ae5d:       eb 81                   jmp    48ade0 <main.TestLogPrintf>
  48ae5f:       cc                      int3
*/

//go:noinline
func TestLogPrintf() {
	log.Printf("string: %v int: %v\n", "汉字", 0xEE)
}

/*
# objdump -d main | awk -v RS= '/^[[:xdigit:]]+ <main.TestLogPrintln>/'
000000000048ae60 <main.TestLogPrintln>:
  48ae60:       49 3b 66 10             cmp    0x10(%r14),%rsp
  48ae64:       0f 86 94 00 00 00       jbe    48aefe <main.TestLogPrintln+0x9e>
  48ae6a:       48 83 ec 60             sub    $0x60,%rsp
  48ae6e:       48 89 6c 24 58          mov    %rbp,0x58(%rsp)
  48ae73:       48 8d 6c 24 58          lea    0x58(%rsp),%rbp
  48ae78:       44 0f 11 7c 24 18       movups %xmm15,0x18(%rsp)
  48ae7e:       44 0f 11 7c 24 28       movups %xmm15,0x28(%rsp)
  48ae84:       44 0f 11 7c 24 38       movups %xmm15,0x38(%rsp)
  48ae8a:       44 0f 11 7c 24 48       movups %xmm15,0x48(%rsp)
  48ae90:       48 8d 15 a9 75 00 00    lea    0x75a9(%rip),%rdx        # 492440 <type.*+0x7440>
  48ae97:       48 89 54 24 18          mov    %rdx,0x18(%rsp)
  48ae9c:       48 8d 35 65 79 03 00    lea    0x37965(%rip),%rsi        # 4c2808 <runtime.defaultGOROOT.str+0x80>
  48aea3:       48 89 74 24 20          mov    %rsi,0x20(%rsp)
  48aea8:       48 89 54 24 28          mov    %rdx,0x28(%rsp)
  48aead:       48 8d 35 44 79 03 00    lea    0x37944(%rip),%rsi        # 4c27f8 <runtime.defaultGOROOT.str+0x70>
  48aeb4:       48 89 74 24 30          mov    %rsi,0x30(%rsp)
  48aeb9:       48 89 54 24 38          mov    %rdx,0x38(%rsp)
  48aebe:       48 8d 15 53 79 03 00    lea    0x37953(%rip),%rdx        # 4c2818 <runtime.defaultGOROOT.str+0x90>
  48aec5:       48 89 54 24 40          mov    %rdx,0x40(%rsp)
  48aeca:       48 8d 15 ef 6e 00 00    lea    0x6eef(%rip),%rdx        # 491dc0 <type.*+0x6dc0>
  48aed1:       48 89 54 24 48          mov    %rdx,0x48(%rsp)
  48aed6:       48 8d 15 5b 78 03 00    lea    0x3785b(%rip),%rdx        # 4c2738 <$f64.bfe62e42fefa39ef+0x18>
  48aedd:       48 89 54 24 50          mov    %rdx,0x50(%rsp)
  48aee2:       48 8d 44 24 18          lea    0x18(%rsp),%rax
  48aee7:       bb 04 00 00 00          mov    $0x4,%ebx
  48aeec:       48 89 d9                mov    %rbx,%rcx
  48aeef:       e8 8c fd ff ff          call   48ac80 <log.Println>
  48aef4:       48 8b 6c 24 58          mov    0x58(%rsp),%rbp
  48aef9:       48 83 c4 60             add    $0x60,%rsp
  48aefd:       c3                      ret
  48aefe:       66 90                   xchg   %ax,%ax
  48af00:       e8 9b 07 fd ff          call   45b6a0 <runtime.morestack_noctxt.abi0>
  48af05:       e9 56 ff ff ff          jmp    48ae60 <main.TestLogPrintln>
  48af0a:       cc                      int3
*/

//go:noinline
func TestLogPrintln() {
	log.Println("string:", "汉字", "int:", 0xFF)
}

func main() {
	TestLogPrintf()
	TestLogPrintln()
}
