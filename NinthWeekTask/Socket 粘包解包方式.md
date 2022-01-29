#### Socket粘包的解包方案

1. ***FixedLengthFrameDecoder***

   固定长度解码器，它能够按照指定的长度对消息进行自动解码，如果接收到的数据不满足设置的固定长度，将等待新的数据到达。

2. ***DelimiterBasedFrameDecoder**

   基于特殊符号的编码器，以自定义分隔符作为包结束标志。

3. ***LengthFieldBasedFrameDecoder***

   自定义协议解码器

   |   **maxFrameLength**    |            **数据包的最大长度**            |
   | :---------------------: | :----------------------------------------: |
   |  **lengthFieldOffset**  | **协议中存储消息长度字段在消息中的偏移量** |
   |  **lengthFieldLength**  |             **长度字段的长度**             |
   |  **lengthAdjustment**   |       **要添加到长度字段值的补偿值**       |
   | **initialBytesToStrip** |            **舍弃的字节的数量**            |

   