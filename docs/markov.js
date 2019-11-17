async function run(data) {
	console.log(wEcho(data));
}

async function jCreateChain(order) {
	throw new Error("sample error");
	if (wNewChain(order) < 0) {
		throw new Error("Error creating chain with order: " + order);
	}
}

async function jReadFile(file) {
	let data = file.files[0];
	const text = await data.text()
	return jRead(text);
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

