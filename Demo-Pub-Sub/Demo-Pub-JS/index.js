const  {connect, StringCodec} = require('nats');

const subject = "my_subject";
const server = 'localhost:4222';

async function pub(){
    const codec = StringCodec();
    const nc =  await connect({server});
    nc.publish(subject,codec.encode('Hello NATS! This is my react publisher, thi is test project'));
    console.log('Message Published Successfully');
    await nc.drain();
}

pub();