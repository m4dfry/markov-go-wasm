
async function run(data) {
	console.log(wEcho(data));
}

async function jCreateChain(order) {
	if (wNewChain(order) < 0) {
		return new Error("Error creating chain with order: " + order);
	}
}

async function jReadFile(file) {
	let data = file.files[0];
	
	data.text().then(text => {
		const chain = wLoad(text);
	});
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

