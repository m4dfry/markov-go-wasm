async function run(data) {
	console.log(wEcho(data));
}

async function jCreateChain(order) {
	if (wNewChain(order) < 0) {
		throw new Error("Error creating chain with order: " + order);
	}
}

async function jRead(text) {
	return wLoad(text);
}

async function jGetJSONChain() {
	const chain = wGetJsonChain();
	if(chain != 0) {
		console.log(JSON.stringify(chain))
	}
}

