#!/bin/sh


/chat --config=/config/chat.yaml &
/rest --config=/config/rest.yaml &
/sess --config=/config/sess.yaml &
/conn --config=/config/conn.yaml &
/task --config=/config/task.yaml &
/demo --config=/config/demo.yaml &


tail -f /dev/null
