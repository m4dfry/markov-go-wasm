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

function jGenerate(prefix, length) {
	const res = wGenerate(prefix, length);
	const genArr = res.split("|");

	if (genArr.length <2 ) {
		throw new Error("Error on WASM generating output.");
	}
	if (genArr[0] != "1") {
		throw new Error("Error generating text:" + genArr[1]);		
	} 
	
	return genArr[1];	
}


