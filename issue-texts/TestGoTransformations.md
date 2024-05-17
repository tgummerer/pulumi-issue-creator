```
=== Failed
=== FAIL: integration/transforms TestGoTransformations/go/simple (32.94s)
    program.go:1946: sample: go/simple
    program.go:1951: pulumi: /home/runner/work/pulumi/pulumi/bin/pulumi
    command.go:46: **** Invoke '/opt/hostedtoolcache/go/1.22.2/x64/bin/go mod edit -replace github.com/pulumi/pulumi/sdk/v3=/home/runner/work/pulumi/pulumi/sdk' in '/home/runner/go/src/stackName-1715764690361049666-427746'
    command.go:121: Command completed without output
    command.go:46: **** Invoke '/opt/hostedtoolcache/go/1.22.2/x64/bin/go mod tidy' in '/home/runner/go/src/stackName-1715764690361049666-427746'
    command.go:121: Command completed without output
    program.go:2096: projdir: /home/runner/go/src/stackName-1715764690361049666-427746
    program.go:1336: Initializing project (dir /home/runner/go/src/stackName-1715764690361049666-427746; stack p-it-fv-az775-2-simple-7166f9e5)
    command.go:46: **** Invoke '/home/runner/work/pulumi/pulumi/bin/pulumi login --cloud-url file:///tmp/TestGoTransformationsgosimple3066846907/001' in '/home/runner/go/src/stackName-1715764690361049666-427746'
    command.go:118: Wrote output to /home/runner/go/src/stackName-1715764690361049666-427746/command-output/pulumi-login.20240515-091812.1596e.log
    command.go:46: **** Invoke '/home/runner/work/pulumi/pulumi/bin/pulumi stack init p-it-fv-az775-2-simple-7166f9e5' in '/home/runner/go/src/stackName-1715764690361049666-427746'
    command.go:118: Wrote output to /home/runner/go/src/stackName-1715764690361049666-427746/command-output/pulumi-stack-init.20240515-091821.47e06.log
    program.go:1503: Performing primary preview and update
    command.go:46: **** Invoke '/home/runner/work/pulumi/pulumi/bin/pulumi up --non-interactive --yes --skip-preview --event-log /tmp/p-it-fv-az775-2-simple-7166f9e5-events.json' in '/home/runner/go/src/stackName-1715764690361049666-427746'
    command.go:98: Invoke '/home/runner/work/pulumi/pulumi/bin/pulumi up --non-interactive --yes --skip-preview --event-log /tmp/p-it-fv-az775-2-simple-7166f9e5-events.json' failed: exit status 255
Updating (p-it-fv-az775-2-simple-7166f9e5):

@ updating.....
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  fatal error: concurrent map writes
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  goroutine 64 [running]:
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  github.com/pulumi/pulumi/sdk/v3/go/pulumi.(*callbackServer).RegisterCallback(0xc000333d40, 0xc0003112d8)
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  	/home/runner/work/pulumi/pulumi/sdk/go/pulumi/callback.go:68 +0xc5
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  github.com/pulumi/pulumi/sdk/v3/go/pulumi.(*Context).registerTransform(0xc000333a40, 0xc000040680)
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  	/home/runner/work/pulumi/pulumi/sdk/go/pulumi/context.go:518 +0xa9
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  github.com/pulumi/pulumi/sdk/v3/go/pulumi.(*Context).registerResource.func1()
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  	/home/runner/work/pulumi/pulumi/sdk/go/pulumi/context.go:1174 +0x2e5
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  created by github.com/pulumi/pulumi/sdk/v3/go/pulumi.(*Context).registerResource in goroutine 1
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  	/home/runner/work/pulumi/pulumi/sdk/go/pulumi/context.go:1157 +0xaf5
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  goroutine 1 [sync.Cond.Wait]:
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  sync.runtime_notifyListWait(0xc00033ca10, 0x0)
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  	/opt/hostedtoolcache/go/1.22.2/x64/src/runtime/sema.go:569 +0x159
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  sync.(*Cond).Wait(0xc000367260?)
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  	/opt/hostedtoolcache/go/1.22.2/x64/src/sync/cond.go:70 +0x85
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  github.com/pulumi/pulumi/sdk/v3/go/internal.(*WorkGroup).Wait(0xc00001a978)
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  	/home/runner/work/pulumi/pulumi/sdk/go/internal/workgroup.go:43 +0xcc
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  github.com/pulumi/pulumi/sdk/v3/go/pulumi.(*Context).wait(0xc000333a40)
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  	/home/runner/work/pulumi/pulumi/sdk/go/pulumi/context.go:232 +0x3d
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  github.com/pulumi/pulumi/sdk/v3/go/pulumi.RunWithContext(0xc000333a40, 0xed5210)
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  	/home/runner/work/pulumi/pulumi/sdk/go/pulumi/run.go:131 +0x28c
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  github.com/pulumi/pulumi/sdk/v3/go/pulumi.runErrInner(0xed5210, 0xed5738, ***0x0, 0x0, 0x0?***)
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  	/home/runner/work/pulumi/pulumi/sdk/go/pulumi/run.go:98 +0x2a5
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  github.com/pulumi/pulumi/sdk/v3/go/pulumi.Run(0xc000100058?, ***0x0?, 0x15e8d88?, 0xc0000061c0?***)
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  	/home/runner/work/pulumi/pulumi/sdk/go/pulumi/run.go:50 +0x28
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  main.main()
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  	/home/runner/go/src/stackName-1715764690361049666-427746/main.go:70 +0x25
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  goroutine 19 [select]:
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  github.com/golang/glog.(*fileSink).flushDaemon(0x16ce978)
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  	/home/runner/go/pkg/mod/github.com/golang/glog@v1.2.0/glog_file.go:351 +0xb9
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  created by github.com/golang/glog.init.1 in goroutine 1
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  	/home/runner/go/pkg/mod/github.com/golang/glog@v1.2.0/glog_file.go:166 +0x126
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  goroutine 24 [select]:
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  google.golang.org/grpc/internal/grpcsync.(*CallbackSerializer).run(0xc00051a810, ***0xfc3c10, 0xc0004aac30***)
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  	/home/runner/go/pkg/mod/google.golang.org/grpc@v1.63.2/internal/grpcsync/callback_serializer.go:76 +0x115
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  created by google.golang.org/grpc/internal/grpcsync.NewCallbackSerializer in goroutine 1
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  	/home/runner/go/pkg/mod/google.golang.org/grpc@v1.63.2/internal/grpcsync/callback_serializer.go:52 +0x11a
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  goroutine 25 [select]:
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  google.golang.org/grpc/internal/grpcsync.(*CallbackSerializer).run(0xc00051a840, ***0xfc3c10, 0xc0004aac80***)
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  	/home/runner/go/pkg/mod/google.golang.org/grpc@v1.63.2/internal/grpcsync/callback_serializer.go:76 +0x115
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  created by google.golang.org/grpc/internal/grpcsync.NewCallbackSerializer in goroutine 1
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  	/home/runner/go/pkg/mod/google.golang.org/grpc@v1.63.2/internal/grpcsync/callback_serializer.go:52 +0x11a
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  goroutine 26 [select]:
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  google.golang.org/grpc/internal/grpcsync.(*CallbackSerializer).run(0xc00051a870, ***0xfc3c10, 0xc0004aacd0***)
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  	/home/runner/go/pkg/mod/google.golang.org/grpc@v1.63.2/internal/grpcsync/callback_serializer.go:76 +0x115
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  created by google.golang.org/grpc/internal/grpcsync.NewCallbackSerializer in goroutine 1
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  	/home/runner/go/pkg/mod/google.golang.org/grpc@v1.63.2/internal/grpcsync/callback_serializer.go:52 +0x11a
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  goroutine 14 [select]:
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  google.golang.org/grpc/internal/transport.(*http2Server).keepalive(0xc0005b0000)
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  	/home/runner/go/pkg/mod/google.golang.org/grpc@v1.63.2/internal/transport/http2_server.go:1168 +0x205
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  created by google.golang.org/grpc/internal/transport.NewServerTransport in goroutine 12
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  	/home/runner/go/pkg/mod/google.golang.org/grpc@v1.63.2/internal/transport/http2_server.go:356 +0x1985
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  goroutine 35 [select]:
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  google.golang.org/grpc/internal/grpcsync.(*CallbackSerializer).run(0xc0003205d0, ***0xfc3c10, 0xc000312190***)
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  	/home/runner/go/pkg/mod/google.golang.org/grpc@v1.63.2/internal/grpcsync/callback_serializer.go:76 +0x115
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  created by google.golang.org/grpc/internal/grpcsync.NewCallbackSerializer in goroutine 1
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  	/home/runner/go/pkg/mod/google.golang.org/grpc@v1.63.2/internal/grpcsync/callback_serializer.go:52 +0x11a
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  goroutine 36 [select]:
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  google.golang.org/grpc/internal/grpcsync.(*CallbackSerializer).run(0xc000320600, ***0xfc3c10, 0xc0003121e0***)
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  	/home/runner/go/pkg/mod/google.golang.org/grpc@v1.63.2/internal/grpcsync/callback_serializer.go:76 +0x115
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  created by google.golang.org/grpc/internal/grpcsync.NewCallbackSerializer in goroutine 1
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  	/home/runner/go/pkg/mod/google.golang.org/grpc@v1.63.2/internal/grpcsync/callback_serializer.go:52 +0x11a
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  goroutine 37 [select]:
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  google.golang.org/grpc/internal/grpcsync.(*CallbackSerializer).run(0xc000320630, ***0xfc3c10, 0xc000312230***)
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  	/home/runner/go/pkg/mod/google.golang.org/grpc@v1.63.2/internal/grpcsync/callback_serializer.go:76 +0x115
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  created by google.golang.org/grpc/internal/grpcsync.NewCallbackSerializer in goroutine 1
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  	/home/runner/go/pkg/mod/google.golang.org/grpc@v1.63.2/internal/grpcsync/callback_serializer.go:52 +0x11a
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  goroutine 15 [IO wait]:
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  internal/poll.runtime_pollWait(0x7f54884c5b58, 0x72)
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  	/opt/hostedtoolcache/go/1.22.2/x64/src/runtime/netpoll.go:345 +0x85
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  internal/poll.(*pollDesc).wait(0xc0002c4100?, 0xc0005a0000?, 0x0)
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  	/opt/hostedtoolcache/go/1.22.2/x64/src/internal/poll/fd_poll_runtime.go:84 +0x27
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  internal/poll.(*pollDesc).waitRead(...)
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  	/opt/hostedtoolcache/go/1.22.2/x64/src/internal/poll/fd_poll_runtime.go:89
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  internal/poll.(*FD).Read(0xc0002c4100, ***0xc0005a0000, 0x8000, 0x8000***)
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  	/opt/hostedtoolcache/go/1.22.2/x64/src/internal/poll/fd_unix.go:164 +0x27a
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  net.(*netFD).Read(0xc0002c4100, ***0xc0005a0000?, 0x8000?, 0x8000?***)
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  	/opt/hostedtoolcache/go/1.22.2/x64/src/net/fd_posix.go:55 +0x25
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  net.(*conn).Read(0xc0001183f0, ***0xc0005a0000?, 0x0?, 0x0?***)
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  	/opt/hostedtoolcache/go/1.22.2/x64/src/net/net.go:179 +0x45
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  bufio.(*Reader).Read(0xc000432540, ***0xc0000f2200, 0x9, 0xc0000774f0?***)
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  	/opt/hostedtoolcache/go/1.22.2/x64/src/bufio/bufio.go:241 +0x197
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  io.ReadAtLeast(***0xfb9d20, 0xc000432540***, ***0xc0000f2200, 0x9, 0x9***, 0x9)
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  	/opt/hostedtoolcache/go/1.22.2/x64/src/io/io.go:335 +0x90
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  io.ReadFull(...)
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  	/opt/hostedtoolcache/go/1.22.2/x64/src/io/io.go:354
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  golang.org/x/net/http2.readFrameHeader(***0xc0000f2200, 0x9, 0xc0000a21e0?***, ***0xfb9d20?, 0xc000432540?***)
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  	/home/runner/go/pkg/mod/golang.org/x/net@v0.23.0/http2/frame.go:237 +0x65
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  golang.org/x/net/http2.(*Framer).ReadFrame(0xc0000f21c0)
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  	/home/runner/go/pkg/mod/golang.org/x/net@v0.23.0/http2/frame.go:498 +0x85
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  google.golang.org/grpc/internal/transport.(*http2Server).HandleStreams(0xc0005b0000, ***0xfc3bd8, 0xc0004d1290***, 0xc0004d12c0)
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  	/home/runner/go/pkg/mod/google.golang.org/grpc@v1.63.2/internal/transport/http2_server.go:645 +0x15e
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  google.golang.org/grpc.(*Server).serveStreams(0xc000151c00, ***0xfc3798?, 0x1730120?***, ***0xfcebc0, 0xc0005b0000***, ***0xfca8d8?, 0xc0001183f0?***)
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  	/home/runner/go/pkg/mod/google.golang.org/grpc@v1.63.2/server.go:1013 +0x3b6
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  google.golang.org/grpc.(*Server).handleRawConn.func1()
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  	/home/runner/go/pkg/mod/google.golang.org/grpc@v1.63.2/server.go:949 +0x56
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  created by google.golang.org/grpc.(*Server).handleRawConn in goroutine 12
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  	/home/runner/go/pkg/mod/google.golang.org/grpc@v1.63.2/server.go:948 +0x1c6
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  goroutine 7 [IO wait]:
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  internal/poll.runtime_pollWait(0x7f54884c5e40, 0x72)
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  	/opt/hostedtoolcache/go/1.22.2/x64/src/runtime/netpoll.go:345 +0x85
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  internal/poll.(*pollDesc).wait(0xc000334200?, 0xc0000e2000?, 0x0)
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  	/opt/hostedtoolcache/go/1.22.2/x64/src/internal/poll/fd_poll_runtime.go:84 +0x27
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  internal/poll.(*pollDesc).waitRead(...)
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  	/opt/hostedtoolcache/go/1.22.2/x64/src/internal/poll/fd_poll_runtime.go:89
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  internal/poll.(*FD).Read(0xc000334200, ***0xc0000e2000, 0x8000, 0x8000***)
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  	/opt/hostedtoolcache/go/1.22.2/x64/src/internal/poll/fd_unix.go:164 +0x27a
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  net.(*netFD).Read(0xc000334200, ***0xc0000e2000?, 0x815f04?, 0xc0000b80a0?***)
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  	/opt/hostedtoolcache/go/1.22.2/x64/src/net/fd_posix.go:55 +0x25
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  net.(*conn).Read(0xc00007c0f8, ***0xc0000e2000?, 0x0?, 0xc0001887b0?***)
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  	/opt/hostedtoolcache/go/1.22.2/x64/src/net/net.go:179 +0x45
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  bufio.(*Reader).Read(0xc0000a22a0, ***0xc0000f2040, 0x9, 0xc?***)
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  	/opt/hostedtoolcache/go/1.22.2/x64/src/bufio/bufio.go:241 +0x197
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  io.ReadAtLeast(***0xfb9d20, 0xc0000a22a0***, ***0xc0000f2040, 0x9, 0x9***, 0x9)
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  	/opt/hostedtoolcache/go/1.22.2/x64/src/io/io.go:335 +0x90
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  io.ReadFull(...)
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  	/opt/hostedtoolcache/go/1.22.2/x64/src/io/io.go:354
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  golang.org/x/net/http2.readFrameHeader(***0xc0000f2040, 0x9, 0xc000188780?***, ***0xfb9d20?, 0xc0000a22a0?***)
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  	/home/runner/go/pkg/mod/golang.org/x/net@v0.23.0/http2/frame.go:237 +0x65
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  golang.org/x/net/http2.(*Framer).ReadFrame(0xc0000f2000)
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  	/home/runner/go/pkg/mod/golang.org/x/net@v0.23.0/http2/frame.go:498 +0x85
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  google.golang.org/grpc/internal/transport.(*http2Client).reader(0xc0000c2488, 0xc0000a2300)
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  	/home/runner/go/pkg/mod/google.golang.org/grpc@v1.63.2/internal/transport/http2_client.go:1602 +0x22d
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  created by google.golang.org/grpc/internal/transport.newHTTP2Client in goroutine 34
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  	/home/runner/go/pkg/mod/google.golang.org/grpc@v1.63.2/internal/transport/http2_client.go:409 +0x1e79
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  goroutine 8 [runnable]:
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  syscall.Syscall(0x1, 0x3, 0xc0000ea000, 0x2f)
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  	/opt/hostedtoolcache/go/1.22.2/x64/src/syscall/syscall_linux.go:69 +0x25
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  syscall.write(0xc000334200?, ***0xc0000ea000?, 0x440600?, 0x1?***)
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  	/opt/hostedtoolcache/go/1.22.2/x64/src/syscall/zsyscall_linux_amd64.go:964 +0x3b
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  syscall.Write(...)
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  	/opt/hostedtoolcache/go/1.22.2/x64/src/syscall/syscall_unix.go:209
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  internal/poll.ignoringEINTRIO(...)
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  	/opt/hostedtoolcache/go/1.22.2/x64/src/internal/poll/fd_unix.go:736
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  internal/poll.(*FD).Write(0xc000334200, ***0xc0000ea000, 0x2f, 0x8000***)
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  	/opt/hostedtoolcache/go/1.22.2/x64/src/internal/poll/fd_unix.go:380 +0x368
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  net.(*netFD).Write(0xc000334200, ***0xc0000ea000?, 0xc0003c6000?, 0xc000597e90?***)
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  	/opt/hostedtoolcache/go/1.22.2/x64/src/net/fd_posix.go:96 +0x25
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  net.(*conn).Write(0xc00007c0f8, ***0xc0000ea000?, 0xc0003c61c0?, 0xc000597ec0?***)
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  	/opt/hostedtoolcache/go/1.22.2/x64/src/net/net.go:191 +0x45
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  google.golang.org/grpc/internal/transport.(*bufWriter).flushKeepBuffer(0xc0000ac0f0)
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  	/home/runner/go/pkg/mod/google.golang.org/grpc@v1.63.2/internal/transport/http_util.go:362 +0x56
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  google.golang.org/grpc/internal/transport.(*bufWriter).Flush(0xc0000ac0f0)
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  	/home/runner/go/pkg/mod/google.golang.org/grpc@v1.63.2/internal/transport/http_util.go:345 +0x1c
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  google.golang.org/grpc/internal/transport.(*loopyWriter).run(0xc0000d2070)
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  	/home/runner/go/pkg/mod/google.golang.org/grpc@v1.63.2/internal/transport/controlbuf.go:591 +0x70
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  google.golang.org/grpc/internal/transport.newHTTP2Client.func6()
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  	/home/runner/go/pkg/mod/google.golang.org/grpc@v1.63.2/internal/transport/http2_client.go:463 +0x85
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  created by google.golang.org/grpc/internal/transport.newHTTP2Client in goroutine 34
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  	/home/runner/go/pkg/mod/google.golang.org/grpc@v1.63.2/internal/transport/http2_client.go:461 +0x242b
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  goroutine 13 [select]:
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  google.golang.org/grpc/internal/transport.(*controlBuffer).get(0xc0004aa1e0, 0x1)
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  	/home/runner/go/pkg/mod/google.golang.org/grpc@v1.63.2/internal/transport/controlbuf.go:418 +0x113
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  google.golang.org/grpc/internal/transport.(*loopyWriter).run(0xc0003c6230)
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  	/home/runner/go/pkg/mod/google.golang.org/grpc@v1.63.2/internal/transport/controlbuf.go:551 +0x86
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  google.golang.org/grpc/internal/transport.NewServerTransport.func2()
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  	/home/runner/go/pkg/mod/google.golang.org/grpc@v1.63.2/internal/transport/http2_server.go:335 +0xe9
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  created by google.golang.org/grpc/internal/transport.NewServerTransport in goroutine 12
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  	/home/runner/go/pkg/mod/google.golang.org/grpc@v1.63.2/internal/transport/http2_server.go:332 +0x193e
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  goroutine 10 [IO wait]:
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  internal/poll.runtime_pollWait(0x7f54884c5d48, 0x72)
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  	/opt/hostedtoolcache/go/1.22.2/x64/src/runtime/netpoll.go:345 +0x85
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  internal/poll.(*pollDesc).wait(0xc000334380?, 0xc0000f8000?, 0x0)
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  	/opt/hostedtoolcache/go/1.22.2/x64/src/internal/poll/fd_poll_runtime.go:84 +0x27
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  internal/poll.(*pollDesc).waitRead(...)
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  	/opt/hostedtoolcache/go/1.22.2/x64/src/internal/poll/fd_poll_runtime.go:89
    pulumi:pulumi:Stack transformations_simple-p-it-fv-az775-2-simple-7166f9e5  internal/poll.(*FD).Read(0xc000334380, ***0xc0000f8000, 0x8000, 0x8000***)
Traceback (most recent call last):
  File "/home/runner/work/pulumi/pulumi/scripts/go-test.py", line 104, in <module>
    raise e
  File "/home/runner/work/pulumi/pulumi/scripts/go-test.py", line 100, in <module>
    sp.check_call(args, shell=False)
  File "/opt/hostedtoolcache/Python/3.12.3/x64/lib/python3.12/subprocess.py", line 413, in check_call
    raise CalledProcessError(retcode, cmd)

[...]

=== FAIL: integration/transforms TestGoTransformations (0.00s)
```

From https://github.com/pulumi/pulumi/actions/runs/9092563159/job/24989923686.  I cut the text up a little bit here to avoid the issue getting too big.

cc @Frassle, do you mind taking a look since it's related to transforms?
