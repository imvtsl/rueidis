// Code generated DO NOT EDIT

package cmds

import "testing"

func server0(s Builder) {
	s.AclCat().Categoryname("1").Build()
	s.AclCat().Build()
	s.AclDeluser().Username("1").Username("1").Build()
	s.AclDryrun().Username("1").Command("1").Arg("1").Arg("1").Build()
	s.AclDryrun().Username("1").Command("1").Build()
	s.AclGenpass().Bits(1).Build()
	s.AclGenpass().Build()
	s.AclGetuser().Username("1").Build()
	s.AclHelp().Build()
	s.AclList().Build()
	s.AclLoad().Build()
	s.AclLog().Count(1).Reset().Build()
	s.AclLog().Count(1).Build()
	s.AclLog().Reset().Build()
	s.AclSave().Build()
	s.AclSetuser().Username("1").Rule("1").Rule("1").Build()
	s.AclSetuser().Username("1").Build()
	s.AclUsers().Build()
	s.AclWhoami().Build()
	s.Bgrewriteaof().Build()
	s.Bgsave().Schedule().Build()
	s.Bgsave().Build()
	s.Command().Build()
	s.CommandCount().Build()
	s.CommandDocs().CommandName("1").CommandName("1").Build()
	s.CommandDocs().Build()
	s.CommandGetkeys().Command("1").Arg("1").Arg("1").Build()
	s.CommandGetkeys().Command("1").Build()
	s.CommandGetkeysandflags().Command("1").Arg("1").Arg("1").Build()
	s.CommandGetkeysandflags().Command("1").Build()
	s.CommandInfo().CommandName("1").CommandName("1").Build()
	s.CommandInfo().Build()
	s.CommandList().FilterbyModuleName("1").Build()
	s.CommandList().FilterbyAclcatCategory("1").Build()
	s.CommandList().FilterbyPatternPattern("1").Build()
	s.CommandList().Build()
	s.ConfigGet().Parameter("1").Parameter("1").Build()
	s.ConfigResetstat().Build()
	s.ConfigRewrite().Build()
	s.ConfigSet().ParameterValue().ParameterValue("1", "1").ParameterValue("1", "1").Build()
	s.Dbsize().Build()
	s.DebugObject().Key("1").Build()
	s.DebugSegfault().Build()
	s.Failover().To().Host("1").Port(1).Force().Abort().Timeout(1).Build()
	s.Failover().To().Host("1").Port(1).Force().Abort().Build()
	s.Failover().To().Host("1").Port(1).Force().Timeout(1).Build()
	s.Failover().To().Host("1").Port(1).Force().Build()
	s.Failover().To().Host("1").Port(1).Abort().Build()
	s.Failover().To().Host("1").Port(1).Timeout(1).Build()
	s.Failover().To().Host("1").Port(1).Build()
	s.Failover().Abort().Build()
	s.Failover().Timeout(1).Build()
	s.Failover().Build()
	s.Flushall().Async().Build()
	s.Flushall().Sync().Build()
	s.Flushall().Build()
	s.Flushdb().Async().Build()
	s.Flushdb().Sync().Build()
	s.Flushdb().Build()
	s.Info().Section("1").Section("1").Build()
	s.Info().Build()
	s.Lastsave().Build()
	s.LatencyDoctor().Build()
	s.LatencyGraph().Event("1").Build()
	s.LatencyHelp().Build()
	s.LatencyHistogram().Command("1").Command("1").Build()
	s.LatencyHistogram().Build()
	s.LatencyHistory().Event("1").Build()
	s.LatencyLatest().Build()
	s.LatencyReset().Event("1").Event("1").Build()
	s.LatencyReset().Build()
	s.Lolwut().Version(1).Build()
	s.Lolwut().Build()
	s.MemoryDoctor().Build()
	s.MemoryHelp().Build()
	s.MemoryMallocStats().Build()
	s.MemoryPurge().Build()
	s.MemoryStats().Build()
	s.MemoryUsage().Key("1").Samples(1).Build()
	s.MemoryUsage().Key("1").Build()
	s.ModuleList().Build()
	s.ModuleLoad().Path("1").Arg("1").Arg("1").Build()
	s.ModuleLoad().Path("1").Build()
	s.ModuleLoadex().Path("1").Config().Config("1", "1").Config("1", "1").Args("1").Args("1").Build()
	s.ModuleLoadex().Path("1").Config().Config("1", "1").Config("1", "1").Build()
	s.ModuleLoadex().Path("1").Args("1").Args("1").Build()
	s.ModuleLoadex().Path("1").Build()
	s.ModuleUnload().Name("1").Build()
	s.Monitor().Build()
	s.Psync().Replicationid("1").Offset(1).Build()
	s.Replicaof().Host("1").Port(1).Build()
	s.Role().Build()
	s.Save().Build()
	s.Shutdown().Nosave().Now().Force().Abort().Build()
	s.Shutdown().Nosave().Now().Force().Build()
	s.Shutdown().Nosave().Now().Abort().Build()
	s.Shutdown().Nosave().Now().Build()
	s.Shutdown().Nosave().Force().Build()
	s.Shutdown().Nosave().Abort().Build()
	s.Shutdown().Nosave().Build()
}

func server1(s Builder) {
	s.Shutdown().Save().Now().Build()
	s.Shutdown().Save().Force().Build()
	s.Shutdown().Save().Abort().Build()
	s.Shutdown().Save().Build()
	s.Shutdown().Now().Build()
	s.Shutdown().Force().Build()
	s.Shutdown().Abort().Build()
	s.Shutdown().Build()
	s.Slaveof().Host("1").Port(1).Build()
	s.SlowlogGet().Count(1).Build()
	s.SlowlogGet().Build()
	s.SlowlogHelp().Build()
	s.SlowlogLen().Build()
	s.SlowlogReset().Build()
	s.Swapdb().Index1(1).Index2(1).Build()
	s.Sync().Build()
	s.Time().Build()
}

func TestCommand_InitSlot_server(t *testing.T) {
	var s = NewBuilder(InitSlot)
	t.Run("0", func(t *testing.T) { server0(s) })
	t.Run("1", func(t *testing.T) { server1(s) })
}

func TestCommand_NoSlot_server(t *testing.T) {
	var s = NewBuilder(NoSlot)
	t.Run("0", func(t *testing.T) { server0(s) })
	t.Run("1", func(t *testing.T) { server1(s) })
}
