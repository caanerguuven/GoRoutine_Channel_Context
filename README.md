# GoRoutine_Channel_Context
I prepared a project that uses go routines, channel and context for async processes.

In this project, Automations run each 15 seconds or in any duration that you want. 
-for that, I used go routines and channel.

In addition, for each automation, 3 accounts run in every 5 seconds until automation duration finishes. (15 second is in my scenario.) 
-for that, I used context structure with routines and channels.

All processes(automation and accounts) are running as async.
