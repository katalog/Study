# Lab 12 Character Sequence Softmax only
import tensorflow as tf
import numpy as np
tf.set_random_seed(777)  # reproducibility

sentence = ("if you want to build a ship, don't drum up people together to "
           "collect wood and don't assign them tasks and work, but rather "
           "teach them to long for the endless immensity of the sea.")

char_set = list(set(sentence))
char_dic = {w:i for i, w in enumerate(char_set)}

dataX = []
dataY = []

data_dim = len(char_set)
hidden_size = len(char_set)
num_classes = len(char_set)
seq_length = 10
learning_rate = 0.1

for i in range(0, len(sentence) - seq_length):
    x_str = sentence[i:i+seq_length]
    y_str = sentence[i+1:i+seq_length+1]
    print (i, x_str, '->', y_str)

    x = [char_dic[c] for c in x_str]
    y = [char_dic[c] for c in y_str]

    dataX.append(x)
    dataY.append(y)

batch_size = len(dataX)

X = tf.placeholder(tf.int32, [None, seq_length])  # X data
Y = tf.placeholder(tf.int32, [None, seq_length])  # Y label

# flatten the data (ignore batches for now). No effect if the batch size is 1
X_one_hot = tf.one_hot(X, num_classes)  # one hot: 1 -> 0 1 0 0 0 0 0 0 0 0


cell = tf.contrib.rnn.BasicLSTMCell(num_units=hidden_size, state_is_tuple=True)
initial_state = cell.zero_state(batch_size, tf.float32)
outputs, _state = tf.nn.dynamic_rnn(cell, X_one_hot, initial_state=initial_state, dtype=tf.float32)

weights = tf.ones([batch_size, seq_length])
sequence_loss = tf.contrib.seq2seq.sequence_loss(
    logits=outputs, targets=Y, weights=weights)
loss = tf.reduce_mean(sequence_loss)  # mean all sequence loss
train = tf.train.AdamOptimizer(learning_rate=learning_rate).minimize(loss)

prediction = tf.argmax(outputs, axis=2)

with tf.Session() as sess:
    sess.run(tf.global_variables_initializer())
    for i in range(7000):
        l, _ = sess.run([loss, train], feed_dict={X: dataX, Y: dataY})
        result = sess.run(prediction, feed_dict={X: dataX})

        # print char using dic

        if i % 333 == 0:
            print(i, "loss:", l)
            for j in range(5):
                org_str = [char_set[c] for c in np.squeeze(dataY[seq_length*j])]
                result_str = [char_set[c] for c in np.squeeze(result[seq_length*j])]
                print(''.join(org_str), " -> ",  ''.join(result_str))
            print('')
        #result_str = [char_set[c] for c in np.squeeze(result)]
        #print(i, "loss:", l, "Prediction:", ''.join(result_str))
        #print(i, "loss:", l, "result:", result)
