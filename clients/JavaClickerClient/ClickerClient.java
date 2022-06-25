package clients.JavaClickerClient;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStream;
import java.io.InputStreamReader;
import java.io.OutputStream;
import java.net.Socket;
import java.net.SocketException;
import java.net.UnknownHostException;

public class ClickerClient {
    private Socket socket;
    private OutputStream writer;
    private InputStream reader;

    public void connect() throws UnknownHostException, IOException {
        this.socket = new Socket("127.0.0.1", 3778);
        this.writer = this.socket.getOutputStream();
        this.reader = this.socket.getInputStream();
    }

    public void connect(String host, int port) throws UnknownHostException, IOException {
        this.socket = new Socket(host, port);
        this.writer = this.socket.getOutputStream();
        this.reader = this.socket.getInputStream();
    }

    public void close() throws SocketException, IOException {
        this.socket.close();
    }

    public String query(String q) throws IOException {
        if (!q.endsWith(";"))
            q += ";";
        this.writer.write(q.getBytes());

        BufferedReader reader = new BufferedReader(new InputStreamReader(this.reader));
        return reader.readLine();
    }
}