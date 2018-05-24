# MBT/CPN: Model-based Testing with Coloured Petri Nets

MBT/CPN is a tool for model-based testing with Coloured Petri Nets. The tool is used as a library on top of CPN Tools and requires installation of CPN Tools. The instructions below assumes that you are familiar with how to use CPN Tools.

## Download

MBT/CPN is distributed as a zip-file containing a set of Standard ML files that needs to be loaded into CPN Tools. Start by download the most recent version provided via this repository.

## Install

Unpack the distribution zip-file in some folder on your machine.

## Use

The examples folder in the zip-file contains an example CPN model twophasecommit.cpn that illustrates how to use MBT/CPN. The example can be found in the examples/tpc folder in the zip-file. To use the example you need to do the following steps:

1. Load the twophasecommit.cpn model into CPN Tools
2. Edit the declaration: `val mbtcpnlibpath =  "c:/work/git/mbtcpnlib/";` to match the folder in which you unpacked the MBT/CPN library
3. Enter the state space tool of CPN Tools
4. Open the MBT page/module and load the library by using ML Evaluate on the ML `use`-statements to load the MBT/CPN library.
5. Use `Execute.ss` and `Execute.sim <n>` to generate test cases using the state space-based approach and the simulation-based approach, respectively. 
6. Use `Execute.export ()` to export the test cases into XML files. These will be put in the output folder in the examples/tpc folder.

To use MBT/CPN in your own model, you need to define the `InEvent` and `OutEvent` colour sets to match the specific CPN model. The file `tcg.sml` in `examples/tpc` gives an example of how to implement the test case specification required by MBT/CPN to generate test cases.
