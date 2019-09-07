
let store = {
    data: {
        txs: [],
        stakeTxId: '',
        balance: 100,
        id: '',
        uid: '',
        privkey: '',
        session_id: '',
        address: '',
        stake: 20,
        email: '',
        coupon: '',
        success: false,
        tx: '',
    },
    setSession(session_id, address) {
        console.log('set sessionid and address', session_id, address)
        this.data.address = address;
        this.data.session_id = session_id;
    },
    setPayout(tx, success) {
        console.log('set tx and success', tx, success)
        this.data.tx = tx;
        this.data.success = success;
    },
    addTx(tx) {
        this.data.txs.push(tx);
        if (tx.action === "Register") {
            this.data.id= tx.id;
            this.data.uid= tx.uid;
            this.data.privkey = tx.privkey;
            this.data.address= tx.address;
            this.data.balance += tx.tokenChange;
        } else if (tx.action === "Stake") {
            this.data.stakeTxId = tx.txid;
            this.data.balance += tx.tokenChange;
        } else if (tx.action === "CompleteLevel") {
            this.data.balance += tx.tokenChange;
        }
    },
    getStakeTxId() {
        return this.data.stakeTxId;
    },
    getMultiplier() {
        return this.data.stake / 20;
    }
};

export default store;
