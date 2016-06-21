%%%----------------------------------------------------------------------
%%% File    : .erl
%%% Author  : LittleTwoLee <bruce>
%%% Purpose :  
%%% Created :  by LittleTwoLee <bruce>
%%%----------------------------------------------------------------------

-module(erlangtemplate).

-author('bruce').

-behaviour(gen_mod).

-include().

-export([start/2, stop/1]).
start(_Host, _Opts) ->
    ok.

stop(_Host) ->
    ok.
