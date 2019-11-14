
async function run(data) {

	console.log(wRead(data));
}

async function jCreateChain(order) {
	console.log(wNewChain(order));
}

async function jReadFile(file) {
	let data = file.files[0];
	
	data.text().then(text => {
		const chain = wLoad(text);
		
	});
}

async function jRead(data) {
	const ret = wLoad(text);
}

async function jGetJSONChain() {
	const ret = wGetJsonChain();
	if(ret != 0) {
		console.log(JSON.stringify(chain))
	}
}

