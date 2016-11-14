# Exercissus

For several weeks,
I've tried to find an easy-to-use HTTP API performance testing tool.
Unfortunately,
I could not find such a tool which simultaneously met my requirements.
I quickly developed Exercissus to serve as a prototype performance test engine,
intending to ask how to accomplish similar tasks with tools already in use at the company (e.g., Grinder or Gatling).
Exercissus is particularly optimized for testing the performance of both stateless and stateful HTTP-hosted APIs.

While you *can* use Exercissus for testing normal web sites or otherwise static web resources,
it's really not built for this purpose.
I'd recommend using another tool (such as those I've surveyed; see next section below).


## Personal Requirements

My requirements for a performance testing tool include,
but aren't necessarily limited to,
what I list below.

* It **must** be trivial to deploy.  Ideally, just copying files onto an ephemeral testing environment should be all that's needed.
* It **must** be ***fast*** to write new scenarios, even for those relatively new to Go.
* It **must** have relatively rapid start-up times.  Long start-up times increases latency to desired results.
* It **must** be memory frugal.  Exercissus may run on relatively low-end OpenStack nodes, which implies that competing, large memory consumption services on the same physical machine will cause thrashing.  This increases run-time delays, not just for Exercissus, but for **all** other services running on the machine *regardless* of which VM it's running in.  By keeping my own memory requirements lean, we reduce our contribution to the problem for other applications running on the physical machine.  Perhaps more importantly, a small memory footprint *reduces the impact* of other high-consumption applications on Exercissus itself while running on the same physical hardware.

## Survey Results of Other Tools
Tool | Requirements
---- | ------------
`The Grinder -tv -series` | Nearly impossible to Google for without the qualifiers indicated.  Java has long start-up times and complicated build and/or deployment.  To be productive with Java, you virtually *need* an IDE of some flavor, which impedes rapid scenario generation.
`ab` or `wrk` | These tools are not optimized for stateful APIs.  Moreover, the scenarios are particularly optimized for hitting a single resource or, if it supports multiple endpoints, it's difficult to use the results of one endpoint to influence how the next is used.  `wrk`, for example, may be scripted in Lua in some cases, but in general, I've not found any documentation online that supports using it in a more general capacity.
Gotling | A Go clone of Gatling which actually looks to be *more* difficult to use than Gatling.  This is because it tries hard to be something it's not: Scala.  It uses a YAML-based DSL instead of a Scala-based DSL, but here again, because of how it works, it seems difficult to virtually impossible to use the results of accessing endpoint 1 to influence how we access endpoint 2.
Gatling | Written in Scala, which implies Java, which implies everything I brought up in the entry for `Grinder`.  Additionally, it also implies my team mates need to learn Yet Another Language to maintain the scenarios.


# Credits

Author | Site | Description | Notes
------ | ---- | ----------- | -----
Rackspace, Inc. | http://gophercloud.io|Go bindings for OpenStack development.
Connor, Glenn | http://mrsharpoblunto.github.io/foswig.js/|Github Javascript Project Name Generator|Used to decide upon this project's name.
